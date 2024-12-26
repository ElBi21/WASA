package database

func (db *appdbimpl) NewSession(token string) error {
	_, err := db.c.Exec("INSERT INTO sessions (token) values {$1}", token)
	return err
}
