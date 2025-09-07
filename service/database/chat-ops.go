package database

import (
	"database/sql"
	"errors"
	"fmt"
	"wasatext/service/customstructs"
)

// CreateConversation creates a new chat in the DB. The chat will be created only if all the users exist
func (db *appdbimpl) CreateConversation(private bool, users []string, name string, description string, photo string) (customstructs.Chat, error) {
	// Create instance of struct that will be used
	queryChat := customstructs.Chat{
		IsPrivate:        private,
		Name:             name,
		GroupDescription: description,
		Photo:            photo,
	}

	// Check if users exist. If not, throw error and DON'T create chat
	for _, user := range users {
		_, err := db.GetUserByName(user)

		if err != nil {
			return queryChat, fmt.Errorf("user %s does not exist, could not create chat", user)
		}
	}

	// After checking, create chat
	_, err := db.c.Exec(
		"INSERT INTO Chats (isPrivate, name, description, photo) VALUES (?, ?, ?, ?);",
		private, name, description, photo)

	if err != nil {
		return queryChat,
			fmt.Errorf("[DB] error while creating new user\n (%w)", err)
	}

	// Add users to the chat
	for _, user := range users {
		_ = db.AddUserToGroup(name, user)
		userStruct, _ := db.GetUserByName(user)
		queryChat.Users = append(queryChat.Users, userStruct)
	}

	return queryChat, err
}

// AddUserToGroup adds a given user into a given chat. The callee must make sure that the chat is not private
// (aside from CreateConversation, which will use this function to add the users to any kind of chat)
func (db *appdbimpl) AddUserToGroup(chat string, user string) error {
	_, err := db.c.Exec(
		"INSERT INTO ChatsUsers (chat, user) VALUES (?, ?);",
		chat, user)

	return err
}

// GetConversation returns a chat given its ChatID. It also retrieves the last sent message
func (db *appdbimpl) GetConversation(chatId int) (customstructs.Chat, error) {
	var returnChat customstructs.Chat

	// Get the chat
	err := db.c.QueryRow("SELECT * FROM Chats WHERE id = ?;", chatId).Scan(
		&returnChat.ID, &returnChat.IsPrivate, &returnChat.Name, &returnChat.GroupDescription, &returnChat.Photo)

	if errors.Is(err, sql.ErrNoRows) {
		return returnChat, fmt.Errorf("[DB] chat does not exist\n (%w)", err)
	}

	// Get last message of the chat
	var lastMsg customstructs.Message
	var userID string

	err = db.c.QueryRow("SELECT * FROM Messages WHERE chat = ? ORDER BY timestamp DESC", chatId).Scan(
		&lastMsg.ID, &userID, &lastMsg.Content, &lastMsg.ChatID, &lastMsg.Timestamp, &lastMsg.Photo,
		&lastMsg.Forwarded, &lastMsg.ReplyingTo, &lastMsg.Deleted)

	if !errors.Is(err, sql.ErrNoRows) {
		lastMsg.Sender, _ = db.GetUserByName(userID)
	}

	// Set last sent message
	returnChat.LastSent = lastMsg

	// Retrieve users from chat
	returnChat.Users = db.GetChatUsers(chatId)

	return returnChat, err
}

// GetChatUsers retrieves the users belonging to a chat
func (db *appdbimpl) GetChatUsers(chatId int) []customstructs.User {
	var users []customstructs.User
	rows, err := db.c.Query("SELECT user FROM ChatsUsers WHERE chat = ?;", chatId)

	if errors.Is(err, sql.ErrNoRows) || rows.Err() != nil {
		return users
	}

	// For each row, retrieve user and append to users
	for rows.Next() {
		var userId string
		_ = rows.Scan(&userId)
		user, _ := db.GetUserByName(userId)
		users = append(users, user)
	}

	return users
}
