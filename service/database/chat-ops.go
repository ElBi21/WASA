package database

import (
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
