package database

import "wasatext/service/customstructs"

func (db *appdbimpl) GetSession(user customstructs.User) (string, error) {
	var session string
	err := db.c.QueryRow("SELECT token FROM sessions WHERE user IS 0").Scan(&session)
	return session, err
}
