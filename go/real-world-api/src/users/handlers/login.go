package handlers

import (
	"log"
	"net/http"
	"real-world-api/src/common"
	UserModel "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

// LoginForm godoc
type LoginForm struct {
	User struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	} `form:"user"`
}

// Login godoc
// @Accept  json
// @Produce  json
// @Produce  json
// @Router /users/login [post]
func Login(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		log.Println(form.User.Email)
		log.Println(form.User.Password)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userModel, err := UserModel.GetUserByEmail(form.User.Email)

	if err != nil {
		if err.Error() == "record not found" {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.Error(err).SetType(gin.ErrorTypePrivate)
	}

	if !checkPassword(form.User.Password, userModel.Password) {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	genLoginResponse(c, userModel)
}

func genLoginResponse(c *gin.Context, userModel *UserModel.UserModel) {
	var user = new(UserSchema)
	user.Email = userModel.Email
	user.Username = userModel.Username
	user.Bio = userModel.Bio
	user.Image = userModel.Image

	tokenString, err := common.JWTSign()
	if err != nil {

		c.Error(err).SetType(gin.ErrorTypePrivate)
	} else {
		user.Token = tokenString
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
