package handlers

import (
	"net/http"
	"real-world-api/src/common"
	userModels "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// LoginForm godoc
type LoginForm struct {
	User struct {
		Email    string `form:"email" binding:"required" example:"jojo@jojo.io"`
		Password string `form:"password" binding:"required" example:"jojojojo"`
	} `form:"user"`
}

// Login godoc
// @tags Users
// @Accept  json
// @Produce  json
// @Produce  json
// @param User body handlers.LoginForm true "User"
// @Success 200 {object} handlers.UserWithTokenSchema "answer"
// @Router /users/login [post]
func Login(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON("Email Or Password Are Required"))
		return
	}

	userModel, err := userModels.GetUserByEmail(form.User.Email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON(common.ErrUserNotExist))
			return
		}

		panic(err)
	}

	if !checkPassword(form.User.Password, userModel.Password) {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON(common.ErrNoCorrectPassword))
		return
	}

	genLoginResponse(c, userModel)
}

func genLoginResponse(c *gin.Context, userModel *userModels.UserModel) {
	var user = new(UserWithTokenSchema)
	user.Email = userModel.Email
	user.Username = userModel.Username
	user.Bio = userModel.Bio
	user.Image = userModel.Image

	tokenString, err := getSignWishUserID(userModel.ID)
	if err != nil {
		panic(err)
	} else {
		user.Token = tokenString
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
