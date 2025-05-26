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
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// Read
	GetName() (string, error)
	GetSession(user customstructs.User) (string, error)

	// Write
	SetName(name string) error

	// Misc
	Ping() error
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

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM Users WHERE type='TEXT NOT NULL';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		tableQueries := [7]string{
			`CREATE TABLE "Users" (
				"id" INTEGER PRIMARY KEY NOT NULL,
				"name" TEXT NOT NULL,
				"photo" BLOB,
				"bio" TEXT
			);`,
			`CREATE TABLE IF NOT EXISTS "Messages" (
				"id" INTEGER PRIMARY KEY NOT NULL,
				"sender" integer NOT NULL,
				"content" TEXT NOT NULL,
				"chat" integer NOT NULL,
				"timestamp" TEXT NOT NULL,
				"photo" BLOB,
				"forwarded" INTEGER
			);`,
			`CREATE TABLE IF NOT EXISTS "Chats" (
				"id" INTEGER PRIMARY KEY NOT NULL,
				"isPrivate" INTEGER NOT NULL,
				"name" TEXT NOT NULL,
				"description" TEXT NOT NULL,
				"photo" BLOB
			);`,
			`CREATE TABLE IF NOT EXISTS "ReceivedMessages" (
				"id" INTEGER PRIMARY KEY NOT NULL,
				"message" integer NOT NULL,
				"received_by" integer NOT NULL
			);`,
			`CREATE TABLE IF NOT EXISTS "SeenMessages" (
				"id" INTEGER PRIMARY KEY NOT NULL,
				"message" integer NOT NULL,
				"seen_by" integer NOT NULL
			);`,
			`CREATE TABLE IF NOT EXISTS "ChatsUsers" (
				"id" INTEGER PRIMARY KEY NOT NULL,
				"chat" integer NOT NULL,
				"user" integer NOT NULL
			);`,
			`CREATE TABLE IF NOT EXISTS "Sessions" (
				"id" INTEGER PRIMARY KEY NOT NULL,
				"user" integer NOT NULL
			);`,
		}

		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
