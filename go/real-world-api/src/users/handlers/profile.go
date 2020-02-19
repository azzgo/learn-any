package handlers

import (
	"net/http"
	"real-world-api/src/common"
	userModels "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

// Profile godoc
// @tags Users
// @Accept  json
// @Produce  json
// @Param username path string true "username"
// @Success 200 {object} handlers.ProfileSchema "answer"
// @Router /profiles/{username} [get]
func Profile(c *gin.Context) {
	value, _ := c.Get(common.KeyJwtCurentUser)
	var currentUserModel *userModels.UserModel
	if value != nil {
		currentUserModel = value.(*userModels.UserModel)
	}

	username := c.Param("username")
	if userModel, _ := userModels.GetUserByUsername(username); *userModel == (userModels.UserModel{}) {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON(common.ErrUserNotExist))
	} else {
		profile := new(ProfileSchema)
		profile.Username = userModel.Username
		profile.Bio = userModel.Bio
		profile.Image = userModel.Image
		profile.Following = checkFollowing(currentUserModel, username)

		c.JSON(http.StatusOK, gin.H{"profile": profile})
	}

}
