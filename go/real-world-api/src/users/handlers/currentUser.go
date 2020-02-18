package handlers

import (
	"net/http"
	"real-world-api/src/common"
	userModels "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

// CurrentUser godoc
// @tags Users
// @Accept  json
// @Produce  json
// @Produce  json
// @Router /user [get]
func CurrentUser(c *gin.Context) {
	value, _ := c.Get(common.KeyJwtCurentUser)
	userModel := value.(*userModels.UserModel)

	var user = new(UserSchema)
	user.Email = userModel.Email
	user.Username = userModel.Username
	user.Bio = userModel.Bio
	user.Image = userModel.Image

	c.JSON(http.StatusOK, gin.H{"user": user})
}
