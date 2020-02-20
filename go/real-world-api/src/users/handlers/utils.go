package handlers

import (
	"real-world-api/src/common"
	userModels "real-world-api/src/users/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func checkPassword(pass string, modelPass string) bool {
	return common.Hash(pass) == modelPass
}

func getSignWishUserID(userid uint) (string, error) {
	return common.JWTSign(jwt.MapClaims{
		"userid":   userid,
		"signtime": time.Now().Unix(),
	})
}

func checkFollowing(currentUserModel *userModels.UserModel, username string) bool {
	isFollowing := false
	currentUserID := func() uint {
		if currentUserModel != nil {
			return currentUserModel.ID
		}
		return 0
	}()

	ids, _ := userModels.GetTargetFollowedIDs(username, 0)
	if currentUserID != 0 {
		for _, id := range ids {
			if id == currentUserID {
				isFollowing = true
			}
		}
	}

	return isFollowing
}
