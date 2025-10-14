/*
Package database is the middleware between the app database and the code. All data
(de)serialization (save/load) from a persistent database are handled here. Database
specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted,
connect to it (using the database data source name from config), and then initialize
an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data
source name (add it to the main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"wasatext/service/customstructs"

	"github.com/ucarion/emoji"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// Seen
	GetUserByName(queryUser string) (customstructs.User, error)
	GetSession(user customstructs.User) (string, error)
	GetUserConversations(user string) (chats []customstructs.Chat)
	GetConversation(chatId int) (customstructs.Chat, error)
	GetChatUsers(chatId int) []customstructs.User
	GetMessage(messageID int, forgetRecvSeen bool) (customstructs.Message, error)
	AddReceivedMessage(messageID int, user string) error
	AddSeenMessage(messageID int, user string) error
	GetWhoReceivedMessage(messageID int) ([]customstructs.User, error)
	GetWhoSawMessage(messageID int) ([]customstructs.User, error)
	GetReactions(messageID int) ([]customstructs.Reaction, error)
	GetAllUsers() []customstructs.User

	// Write
	RegisterNewUser(newUserName string) (customstructs.User, error)
	SetNewUserName(user string, newUserName string) error
	SetNewDisplayName(user string, newDisplayName string) error
	SetNewBiography(user string, newBiography string) error
	SetNewUserPhoto(user string, newPhoto string) error
	CreateConversation(private bool, users []string, name string, description string, photo string) (customstructs.Chat, error)
	AddUserToGroup(chat int, user string) error
	CreateMessage(newMessage customstructs.PrimordialMessage) (customstructs.Message, error)
	ForwardMessage(message customstructs.Message) (int, error)
	CreateReaction(message int, content emoji.Emoji, user string) (customstructs.Reaction, error)
	DeleteReaction(reactionID int) error
	DeleteMessage(messageID int) (customstructs.Message, error)
	SetNewGroupDescription(chatID int, newDescription string) error
	RemoveUserFromGroup(chatID int, userID string) error
	DeleteChatAndMessages(chat customstructs.Chat) (error, error)
	SetNewGroupName(chatID int, newName string) error
	SetNewGroupPhoto(chatID int, newPhoto string) error

	// Misc
	Ping() error
	GetLastReactionID() (ID int)
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Start the transaction
	tx, err := db.Begin()

	// Check for transaction errors
	if err != nil {
		return nil, err
	}

	tableQueries := [8]string{
		`CREATE TABLE IF NOT EXISTS "Users" (
			"name" TEXT PRIMARY KEY NOT NULL,
			"display_name" TEXT,
			"photo" BLOB NOT NULL,
			"bio" TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS "Messages" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			"sender" TEXT NOT NULL,
			"content" TEXT NOT NULL,
			"chat" INTEGER NOT NULL,
			"timestamp" DATETIME NOT NULL,
			"photo" BLOB,
			"forwarded" INTEGER NOT NULL,
			"replying_to" INTEGER NOT NULL,
			"deleted" INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS "Chats" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			"isPrivate" INTEGER NOT NULL,
			"name" TEXT NOT NULL,
			"description" TEXT NOT NULL,
			"photo" BLOB
		);`,
		`CREATE TABLE IF NOT EXISTS "ReceivedMessages" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			"message" INTEGER NOT NULL,
			"received_by" TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS "SeenMessages" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			"message" INTEGER NOT NULL,
			"seen_by" TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS "ChatsUsers" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			"chat" INTEGER NOT NULL,
			"user" TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS "Sessions" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			"user" TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS "Reactions" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			"message" INTEGER NOT NULL,
			"content" TEXT NOT NULL,
			"user" TEXT NOT NULL
		);`,
	}

	// For each query, execute it
	for i := range tableQueries {
		_, err := tx.Exec(tableQueries[i])

		// If the query had an issue, return nil
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			return nil,
				fmt.Errorf("error while creating table %d\n (%w)", i+1, err)
		}
	}

	// Commit the transaction and check for errors
	transError := tx.Commit()

	if transError != nil {
		return nil,
			fmt.Errorf("error while committing the DB creation queries\n (%w)", transError)
	}

	// Finally, return the DB
	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
