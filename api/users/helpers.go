package users

import (
	"database/sql"
)

type IdError string

func (e IdError) Error() string {
	return string(e)
}

func checkUsername(username string, db *sql.DB) error {
	row := db.QueryRow("SELECT id FROM users WHERE username = $1", username)
	var id int
	if err := row.Scan(&id); err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	return IdError("users.username_taken")
}

func checkEmail(email string, db *sql.DB) error {
	row := db.QueryRow("SELECT id FROM users WHERE email = $1", email)
	var id int
	if err := row.Scan(&id); err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	return IdError("users.email_taken")
}