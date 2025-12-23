package database

import (
	"database/sql"
	"errors"
	"wasatext/service/customstructs"

	"github.com/ucarion/emoji"
)

// GetLastReactionID is a function that retrieves the last used ID for the reactions
func (db *appdbimpl) GetLastReactionID() (reactionID int) {
	err := db.c.QueryRow(`SELECT id FROM Reactions ORDER BY id DESC`).Scan(&reactionID)

	// If there exists no messages, then this message is the first ever sent
	if errors.Is(err, sql.ErrNoRows) {
		reactionID = 1
	} else {
		reactionID += 1
	}

	return reactionID
}

// GetReactions retrieves any reactions attached to a given message from the DB
func (db *appdbimpl) GetReactions(messageID int) ([]customstructs.Reaction, error) {
	var reactions []customstructs.Reaction

	queryRes, err := db.c.Query("SELECT * FROM Reactions WHERE message = ?", messageID)

	if err != nil || queryRes.Err() != nil {
		return reactions, err
	}

	for queryRes.Next() {
		// Create reaction
		var reaction customstructs.Reaction
		var sender string
		var emojiContent string
		var okEmoji bool

		err = queryRes.Scan(&reaction.ID, &reaction.RefMessage, &emojiContent, &sender)

		if err != nil {
			return reactions, err
		}

		reaction.Content, okEmoji = emoji.Lookup(emojiContent)

		if !okEmoji {
			return reactions, err
		}

		// Retrieve user
		reaction.Sender, err = db.GetUserByName(sender)

		if err != nil {
			return reactions, err
		}

		reactions = append(reactions, reaction)
	}

	_ = queryRes.Close()

	return reactions, nil
}

// CreateReaction creates a reaction given the message, the reaction's content and the sending user
func (db *appdbimpl) CreateReaction(message int, content emoji.Emoji, user string) (customstructs.Reaction, error) {
	ID := db.GetLastReactionID()
	_, err := db.c.Exec("INSERT INTO Reactions (id, message, content, user) VALUES (?, ?, ?, ?)",
		ID, message, content.FullyQualifiesAs, user)

	if err != nil {
		return customstructs.Reaction{}, err
	}

	// Retrieve user object
	var sender customstructs.User
	sender, err = db.GetUserByName(user)

	if err != nil {
		return customstructs.Reaction{}, err
	}

	// Build reaction object
	reaction := customstructs.Reaction{
		ID:         ID,
		RefMessage: message,
		Sender:     sender,
		Content:    content,
	}

	return reaction, nil
}

// DeleteReaction removes a reaction from the database
func (db *appdbimpl) DeleteReaction(reactionID int) error {
	_, err := db.c.Exec("DELETE FROM Reactions WHERE id = ?", reactionID)

	return err
}
