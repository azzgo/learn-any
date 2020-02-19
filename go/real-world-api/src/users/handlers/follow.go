package handlers

import (
	"net/http"
	"real-world-api/src/common"
	userModels "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

// Follow godoc
// @tags Users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Success 200 {object} handlers.ProfileSchema "answer"
// @Router /api/profiles/:username/follow [post]
func Follow(c *gin.Context) {
	value, _ := c.Get(common.KeyJwtCurentUser)
	curUserModel := value.(*userModels.UserModel)

	username := c.Param("username")
	
	targetUserModel, _ := userModels.GetUserByUsername(username)
	err := userModels.FollowUser(curUserModel, username)
	
	if err != nil {
		panic(err)
	}

	var profile = new(ProfileSchema)
	profile.Bio = targetUserModel.Bio
	profile.Image = targetUserModel.Image
	profile.Username = targetUserModel.Username
	profile.Following = true
	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

// UnFollow godoc
// @tags Users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Success 200 {object} handlers.ProfileSchema "answer"
// @Router /api/profiles/:username/follow [post]
func UnFollow(c *gin.Context) {
	value, _ := c.Get(common.KeyJwtCurentUser)
	curUserModel := value.(*userModels.UserModel)

	username := c.Param("username")
	
	targetUserModel, _ := userModels.GetUserByUsername(username)
	err := userModels.UnFollowUser(curUserModel, username)
	
	if err != nil {
		panic(err)
	}

	var profile = new(ProfileSchema)
	profile.Bio = targetUserModel.Bio
	profile.Image = targetUserModel.Image
	profile.Username = targetUserModel.Username
	profile.Following = false

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}
