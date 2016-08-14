package utils

import (
	"log"
	"unicode/utf8"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"github.com/kataras/iris"

	"reactizer-go/config"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

// 'Authorize' checks the 'X-Authorization' header if it contains the JWT token required by some
// queries. If the token is there, it is decoded into a user id and returned.
//
// Errors:
// "auth.no_auth_header"
func Authorize(c *iris.Context) (int, error) {
	token := c.RequestHeader("X-Authorization")
	if token == "" {
		return 0, AuthError("auth.no_auth_header")
	}

	uid, err := DecodeToken(token)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

// 'HashPassword' hashes the given password
func HashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, config.CryptCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// 'VerifyPassword' verifies if the given password and hash match
//
// Errors:
// "auth.invalid_password"
func VerifyPassword(password, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		log.Print(err)
		return AuthError("auth.invalid_password")
	}
	return nil
}

// 'CheckPassword' checks a given password's complexity.
//
// Errors:
// "auth.password_too_short"
// "auth.password_too_long"
// "auth.password_no_number"
// "auth.password_no_upper"
// "auth.password_no_lower"
func CheckPassword(password string) error {
	if utf8.RuneCountInString(password) < 8 {
		return AuthError("auth.password_too_short")
	}
	if utf8.RuneCountInString(password) > 32 {
		return AuthError("auth.password_too_long")
	}
	if match, _ := regexp.MatchString(`\d`, password); !match {
		return AuthError("auth.password_no_number")
	}
	if match, _ := regexp.MatchString(`[A-Z]`, password); !match {
		return AuthError("auth.password_no_upper")
	}
	if match, _ := regexp.MatchString(`[a-z]`, password); !match {
		return AuthError("auth.password_no_lower")
	}
	return nil
}
