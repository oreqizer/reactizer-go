package utils

import (
	"database/sql"
	"log"

	"github.com/kataras/iris"
	"unicode/utf8"
	"regexp"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

// 'Authorize' checks the 'X-Authorization' header if it contains the JWT token required by some
// queries. If the token is there, it is decoded into a user id and returned.
//
// In case of an error, translation id AuthError is returned.
func Authorize(c *iris.Context, db *sql.DB) (int, error) {
	token := c.RequestHeader("X-Authorization")
	if token == "" {
		return 0, AuthError("auth.no_auth_header")
	}

	log.Print(decodeToken(token)) // TODO: create token
	return 0, nil
}

// 'CheckPassword' checks a given password's complexity.
func CheckPassword(password string) error {
	if utf8.RuneCountInString(password) < 8 {
		return AuthError("auth.password_too_short")
	}
	if utf8.RuneCountInString(password) > 32 {
		return AuthError("auth.password_too_long")
	}
	if match, err := regexp.MatchString(`\d`, password); err != nil {
		return err
	} else if !match {
		return AuthError("auth.password_no_number")
	}
	if match, err := regexp.MatchString(`[A-Z]`, password); err != nil {
		return err
	} else if !match {
		return AuthError("auth.password_no_upper")
	}
	if match, err := regexp.MatchString(`[a-z]`, password); err != nil {
		return err
	} else if !match {
		return AuthError("auth.password_no_lower")
	}
	return nil
}
