package users

import "database/sql"

type IdError error

func (e IdError) Error() string {
	return string(e)
}

func checkUsername(username string, db *sql.DB) error {
	row := db.QueryRow("SELECT id FROM users WHERE username = ?", username)
	if err := row.Scan(); err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	return IdError("users.username_taken")
}

func checkEmail(email string, db *sql.DB) error {
	row := db.QueryRow("SELECT id FROM users WHERE email = ?", email)
	if err := row.Scan(); err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	return IdError("users.password_taken")
}