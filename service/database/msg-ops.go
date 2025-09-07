package database

import (
	"database/sql"
	"errors"
	"fmt"
	"wasatext/service/customstructs"
)

// CreateMessage creates a new message
func (db *appdbimpl) CreateMessage(newMessage customstructs.PrimordialMessage) (customstructs.Message, error) {
	var message customstructs.Message
	var ID int

	err := db.c.QueryRow(`SELECT id FROM Messages ORDER BY id DESC`).Scan(&ID)

	// If there exists no messages, then this message is the first ever sent
	if errors.Is(err, sql.ErrNoRows) {
		ID = 1
	} else {
		ID += 1
	}

	userSender, errUser := db.GetUserByName(newMessage.Sender)
	_, errChat := db.GetConversation(newMessage.ChatID)

	// If the sender or the chat do not exist, return an error
	if errUser != nil {
		return message, errUser
	} else if errChat != nil {
		return message, errChat
	}

	// Create message on DB
	_, err = db.c.Exec(
		`INSERT INTO Messages (id, sender, content, chat, timestamp, photo, forwarded, replying_to, deleted)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, ID, newMessage.Sender, newMessage.Content, newMessage.ChatID,
		newMessage.Timestamp, newMessage.Photo, newMessage.Forwarded, newMessage.ReplyingTo, false)

	if err != nil {
		return message, err
	}

	// Switch from PrimordialMessage to Message
	message.ID = ID
	message.Sender = userSender
	message.ChatID = newMessage.ChatID
	message.Timestamp = newMessage.Timestamp
	message.Photo = newMessage.Photo
	message.Forwarded = newMessage.Forwarded
	message.ReplyingTo = newMessage.ReplyingTo

	return message, nil
}

// GetMessage retrieves a message given a messageID
// TODO: Implement retrieval of reactions
func (db *appdbimpl) GetMessage(messageID int, forgetRecvSeen bool) (customstructs.Message, error) {
	var message customstructs.Message
	var userID string

	// Get message from its ID
	res := db.c.QueryRow("SELECT * FROM Messages WHERE id = ?", messageID)

	if res.Err() != nil {
		return message, fmt.Errorf("[DB] there exists no message with id %d: %w", messageID, res.Err())
	}

	// Fill the object
	err := res.Scan(&message.ID, &userID, &message.Content, &message.ChatID,
		&message.Timestamp, &message.Photo, &message.Forwarded, &message.ReplyingTo, &message.Deleted)

	if err != nil {
		return message, err
	}

	// Get the sender of the message from the DB as an object
	message.Sender, err = db.GetUserByName(userID)

	if err != nil {
		return message, err
	}

	// If the flag is active, we don't retrieve who received and saw the message
	if !forgetRecvSeen {
		// Retrieve who received and saw the message
		whoReceived, errRecv := db.GetWhoReceivedMessage(message.ID)
		whoSaw, errSeen := db.GetWhoSawMessage(message.ID)

		if errRecv != nil {
			return message, errRecv
		} else if errSeen != nil {
			return message, errSeen
		}

		// Update arrays
		message.Received = whoReceived
		message.Seen = whoSaw
	}

	return message, nil
}

// AddReceivedMessage marks a message as received by a given user
func (db *appdbimpl) AddReceivedMessage(messageID int, user string) error {
	_, errUser := db.GetUserByName(user)
	_, errChat := db.GetMessage(messageID, true)

	// If either the user or the message do not exist, then abort
	if errUser != nil || errChat != nil {
		return fmt.Errorf("[DB] either the user %s or the message with ChatID %d do not exist", user, messageID)
	}

	// Check if user already received
	res := db.c.QueryRow(`SELECT message, received_by FROM ReceivedMessages
                            WHERE message = ? AND received_by = ?`, messageID, user).Scan()

	if !errors.Is(res, sql.ErrNoRows) {
		return fmt.Errorf("[DB] The user %s has already marked message %d as received", user, messageID)
	}

	// Insert the tuple into the DB
	_, err := db.c.Exec("INSERT INTO ReceivedMessages (message, received_by) VALUES (?, ?)",
		messageID, user)

	return err
}

// AddSeenMessage marks a message as seen by a given user
func (db *appdbimpl) AddSeenMessage(messageID int, user string) error {
	_, errUser := db.GetUserByName(user)
	_, errChat := db.GetMessage(messageID, true)

	// If either the user or the message do not exist, then abort
	if errUser != nil || errChat != nil {
		return fmt.Errorf("[DB] either the user %s or the message with ChatID %d do not exist", user, messageID)
	}

	// Check if user already saw
	res := db.c.QueryRow(`SELECT message, seen_by FROM SeenMessages
                            WHERE message = ? AND seen_by = ?`, messageID, user).Scan()

	if !errors.Is(res, sql.ErrNoRows) {
		return fmt.Errorf("[DB] The user %s has already marked message %d as seen", user, messageID)
	}

	// Insert the tuple into the DB
	_, err := db.c.Exec("INSERT INTO SeenMessages (message, seen_by) VALUES (?, ?)",
		messageID, user)

	return err
}

// ForwardMessage forwards a message by "sending" it in another chat. Mind that the fields in message
// must be already adapted for the forwarding
func (db *appdbimpl) ForwardMessage(message customstructs.Message) (int, error) {
	var ID int
	err := db.c.QueryRow(`SELECT id FROM Messages ORDER BY id DESC`).Scan(&ID)

	// If there exists no messages, then this message is the first ever sent
	if errors.Is(err, sql.ErrNoRows) {
		ID = 1
	} else {
		ID += 1
	}

	_, err = db.c.Exec(
		`INSERT INTO Messages (id, sender, content, chat, timestamp, photo, forwarded, replying_to, deleted)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, ID, message.Sender.Name, message.Content, message.ChatID,
		message.Timestamp, message.Photo, message.Forwarded, message.ReplyingTo, false)

	if err != nil {
		return ID, err
	}

	return ID, nil
}

// GetWhoReceivedMessage checks who received a message, and returns an array of users who actually received it
func (db *appdbimpl) GetWhoReceivedMessage(messageID int) ([]customstructs.User, error) {
	var users []customstructs.User

	queryRes, err := db.c.Query("SELECT received_by FROM ReceivedMessages WHERE message = ?", messageID)

	if errors.Is(err, sql.ErrNoRows) || queryRes.Err() != nil {
		return users, fmt.Errorf("[DB] No rows found, the message with ID %d does not exist", messageID)
	}

	for queryRes.Next() {
		var user customstructs.User
		var userID string
		_ = queryRes.Scan(&userID)

		user, err = db.GetUserByName(userID)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	_ = queryRes.Close()

	return users, nil
}

// GetWhoSawMessage checks who saw a message, and returns an array of users who actually have seen it
func (db *appdbimpl) GetWhoSawMessage(messageID int) ([]customstructs.User, error) {
	var users []customstructs.User

	queryRes, err := db.c.Query("SELECT seen_by FROM SeenMessages WHERE message = ?", messageID)

	if errors.Is(err, sql.ErrNoRows) || queryRes.Err() != nil {
		return users, fmt.Errorf("[DB] No rows found, the message with ID %d does not exist", messageID)
	}

	for queryRes.Next() {
		var user customstructs.User
		var userID string
		_ = queryRes.Scan(&userID)

		user, err = db.GetUserByName(userID)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	_ = queryRes.Close()

	return users, nil
}

// DeleteMessage flags a message as "removed"
func (db *appdbimpl) DeleteMessage(messageID int) error {
	_, err := db.c.Exec("UPDATE Messages SET deleted = 1 WHERE id = ?", messageID)

	return err
}
