package utils

import (
	"log"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"reactizer-go/config"
)


func GetToken(password string, uid int) (string, error) {
	raw := jwt.NewWithClaims(jwt.SigningMethodRS384, jwt.MapClaims{
    "sub": uid,
    "iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 10).Unix(),
	})

	token, err := raw.SignedString(config.Secret)
	if err != nil {
		return "", err
	}
	return token
}

func DecodeToken(raw string) (int, error) {
	token, err := jwt.Parse(raw, keyfunc)
	if err != nil {
		log.Print(err)
		return 0, AuthError("auth.invalid_token")
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    //return claims["sub"].(int), nil
    return 1, nil
	}

	return 0, AuthError("auth.invalid_token")
}

func keyfunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected method: %v", token.Header["alg"])
	}
	return config.Secret, nil
}
