package common

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTSign godoc
func JWTSign() (string, error) {
	secret := Config["SECRET"]

	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString([]byte(secret))
}

// JWTValidate godoc
func JWTValidate(tokenString string) (bool, error) {
	secret := Config["SECRET"]
	token, err := jwt.Parse(tokenString, func(ttToken *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	return token.Valid, err
}
