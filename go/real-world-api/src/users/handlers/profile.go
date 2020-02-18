package handlers

import (
	"net/http"
	"real-world-api/src/common"
	userModels "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

// Profile godoc
func Profile(c *gin.Context) {
	username := c.Param("username")
	if userModel, _ := userModels.GetUserByUsername(username); *userModel == (userModels.UserModel{}) {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON(common.ErrUserNotExist))
	} else {
		profile := new(ProfileSchema)
		profile.Username = userModel.Username
		profile.Bio = userModel.Bio
		profile.Image = userModel.Image
		profile.Following = userModel.Following
		c.JSON(http.StatusOK, gin.H{"profile": profile})
	}

}
