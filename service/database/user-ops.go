package database

import (
	"fmt"
	"wasatext/service/customstructs"
)

// GetUserByName returns from the database all the rows with the given name
func (db *appdbimpl) GetUserByName(queryUser string) (customstructs.User, error) {
	var resultUser customstructs.User

	// Perform the query. If the user exists, it will retrieve it from the DB, else it will return error,
	// meaning that the user must be created
	err := db.c.QueryRow("SELECT * FROM Users WHERE name = ?", queryUser).Scan(
		&resultUser.Name,
		&resultUser.DisplayName,
		&resultUser.ProfilePic,
		&resultUser.Biography)

	return resultUser, err
}

// RegisterNewUser creates a new user in the DB given a name. The display_name will be initialized as name
func (db *appdbimpl) RegisterNewUser(user string) (customstructs.User, error) {
	queryUser := customstructs.User{
		Name:        user,
		DisplayName: user,
		ProfilePic:  "",
		Biography:   "Hey, I'm a WASAText user",
	}

	_, err := db.c.Exec(
		"INSERT INTO Users (name, display_name, photo, bio) VALUES (?, ?, ?, ?)",
		queryUser.Name, queryUser.DisplayName, queryUser.ProfilePic, queryUser.Biography)

	if err != nil {
		return queryUser,
			fmt.Errorf("[DB] error while creating new user\n (%w)", err)
	}

	return queryUser, err
}

// SetNewUserName changes the username of a user, given the user and a new username.
// When this function is called, the callee is sure that the new display name has a length
// that is 3 <= len <= 16
func (db *appdbimpl) SetNewUserName(user string, newUserName string) error {
	_, err := db.c.Exec("UPDATE Users SET name = ? WHERE name = ?",
		newUserName, user)

	if err != nil {
		return fmt.Errorf("[DB] error while changing username\n (%w)", err)
	}

	return nil
}

// SetNewDisplayName changes the display name of a user, given the user and a new display name.
// When this function is called, the callee is sure that the new display name has a length
// that is 3 <= len <= 32
func (db *appdbimpl) SetNewDisplayName(user string, newDisplayName string) error {
	_, err := db.c.Exec("UPDATE Users SET display_name = ? WHERE name = ?",
		newDisplayName, user)

	if err != nil {
		return fmt.Errorf("[DB] error while changing display name\n (%w)", err)
	}

	return nil
}

// SetNewBiography changes the biography of a user, given the user and a new biography.
// The callee must be sure that 1 <= len(bio) <= 512
func (db *appdbimpl) SetNewBiography(user string, newBiography string) error {
	_, err := db.c.Exec("UPDATE Users SET bio = ? WHERE name = ?",
		newBiography, user)

	if err != nil {
		return fmt.Errorf("[DB] error while changing biography\n (%w)", err)
	}

	return nil
}

// SetNewUserPhoto changes the profile picture of a user, given the user and a new biography.
// The base64 encoding must be smaller than 4294967296 characters
func (db *appdbimpl) SetNewUserPhoto(user string, newPhoto string) error {
	_, err := db.c.Exec("UPDATE Users SET photo = ? WHERE name = ?",
		newPhoto, user)

	if err != nil {
		return fmt.Errorf("[DB] error while changing biography\n (%w)", err)
	}

	return nil
}

// GetUserConversations gets all the chats to which a user belongs. Returns a customstructs.Chat struct
func (db *appdbimpl) GetUserConversations(user string) []customstructs.Chat {
	var chats []customstructs.Chat
	queryChats, err := db.c.Query(`WITH user_chats AS (
    SELECT r.id, r.isPrivate, r.name, r.description, r.photo
    FROM ChatsUsers l INNER JOIN Chats r ON l.chat = r.id
    WHERE user = ?
	),
	chats_with_message AS (
		WITH last_messages AS (
			SELECT id, chat, max(timestamp) AS "order_tstamp"
			FROM Messages
			GROUP BY chat
		)
		SELECT r.id, l.order_tstamp, 1 AS custom_order
		FROM last_messages l INNER JOIN user_chats r ON l.chat = r.id
	),
	chats_without_messages AS (
		SELECT id, 0 AS order_tstamp, 2 AS custom_order 
		FROM (
			SELECT id
			FROM user_chats
				EXCEPT
			SELECT id
			FROM chats_with_message
    	)
	)
	SELECT *
	FROM chats_without_messages
		UNION
	SELECT *
	FROM chats_with_message
	ORDER BY custom_order DESC, order_tstamp DESC;`, user)

	// If the user belongs to no chats, then return an empty array
	if err != nil || queryChats.Err() != nil {
		return chats
	}

	// If the user does belong to some chats, close the query and unwrap the result
	defer queryChats.Close()

	for queryChats.Next() {
		var userChat customstructs.Chat
		var chatID int
		_ = queryChats.Scan(&chatID, nil, nil)

		userChat, err = db.GetConversation(chatID)

		if err != nil {
			return chats
		}

		// Append chat to return array
		chats = append(chats, userChat)
	}

	return chats
}
