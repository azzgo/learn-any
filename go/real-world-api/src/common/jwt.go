package common

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// JWTSign godoc
func JWTSign(mapClaims jwt.MapClaims) (string, error) {
	secret := Config["SECRET"]

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(secret))
}

const (
	// KeyJwtCurentUser godoc
	KeyJwtCurentUser = "KeyJwtCurentUser"
)

// JWTValidate godoc
func JWTValidate(tokenString string) (*jwt.Token, error) {
	secret := []byte(Config["SECRET"])
	token, err := jwt.Parse(tokenString, func(ttToken *jwt.Token) (interface{}, error) {
		if _, ok := ttToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", ttToken.Header["alg"])
		}

		return secret, nil
	})

	return token, err
}
