package utils

import (
	"log"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"reactizer-go/config"
)

func GetToken(uid int) (string, error) {
	raw := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
    "sub": uid,
    "iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 10).Unix(),
	})

	token, err := raw.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func DecodeToken(raw string) (int, error) {
	token, err := jwt.Parse(raw, keyfunc)
	if err != nil {
		log.Print(err)
		return 0, AuthError("auth.invalid_token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    return int(claims["sub"].(float64)), nil
	}

	return 0, AuthError("auth.invalid_token")
}

func keyfunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected method: %v", token.Header["alg"])
	}
	return []byte(config.Secret), nil
}
