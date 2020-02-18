package handlers

import (
	"real-world-api/src/common"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func checkPassword(pass string, modelPass string) bool {
	return common.Hash(pass) == modelPass
}

func getSignWishUserID(userid uint) (string, error) {
	return common.JWTSign(jwt.MapClaims{
		"userid": userid,
		"signtime": time.Now().Unix(),
	})
}
