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
func (db *appdbimpl) RegisterNewUser(newUserName string) (customstructs.User, error) {
	queryUser := customstructs.User{
		Name:        newUserName,
		DisplayName: newUserName,
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
