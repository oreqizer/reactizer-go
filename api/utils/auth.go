package utils

import (
	"unicode/utf8"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"github.com/kataras/iris"
	"github.com/golang/glog"

	"reactizer-go/config"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

// 'Authorize' checks the 'X-Authorization' header if it contains the JWT token required by some
// queries. If the token is there, it is decoded into a user id and returned.
func Authorize(c *iris.Context) (int, error) {
	token := c.RequestHeader("X-Authorization")
	if token == "" {
		return 0, AuthError(noAuthHeader)
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
func VerifyPassword(password, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		glog.Error(err)
		return AuthError(invalidPassword)
	}
	return nil
}

// 'CheckPassword' checks a given password's complexity.
func CheckPassword(password string) error {
	if utf8.RuneCountInString(password) < 8 {
		return AuthError(passwordTooShort)
	}
	if utf8.RuneCountInString(password) > 32 {
		return AuthError(passwordTooLong)
	}
	if match, _ := regexp.MatchString(`\d`, password); !match {
		return AuthError(passwordNoNumber)
	}
	if match, _ := regexp.MatchString("[A-Z]", password); !match {
		return AuthError(passwordNoUpper)
	}
	if match, _ := regexp.MatchString("[a-z]", password); !match {
		return AuthError(passwordNoLower)
	}
	return nil
}
