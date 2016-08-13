package utils

import (
	"log"
	"fmt"

	"github.com/dgrijalva/jwt-go"

	"reactizer-go/config"
)

func decodeToken(raw string) (int, error) {
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
