package users

import (
	"database/sql"
	"unicode/utf8"

	"reactizer-go/api/utils"
	"regexp"
)

func checkUsername(username string, db *sql.DB) error {
	row := db.QueryRow("SELECT id FROM users WHERE username = ?", username)
	if err := row.Scan(); err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}

	return ""
}

func checkEmail(email string, db *sql.DB) error {
	row := db.QueryRow("SELECT id FROM users WHERE email = ?", email)
	if err := row.Scan(); err == sql.ErrNoRows {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return false, nil
}

func checkPassword(password string) error {
	if utf8.RuneCountInString(password) < 8 {
		return utils.AuthError("auth.password_too_short")
	}
	if utf8.RuneCountInString(password) > 32 {
		return utils.AuthError("auth.password_too_long")
	}
	if match, err := regexp.MatchString(`\d`, password); err != nil {
		return err
	} else if !match {
		return utils.AuthError("auth.password_no_number")
	}// TODO
	if match, err := regexp.MatchString(`\d`, password); !match || err != nil {
		return utils.AuthError("auth.password_no_number")
	}
	if match, err := regexp.MatchString(`\d`, password); !match || err != nil {
		return utils.AuthError("auth.password_no_number")
	}
}