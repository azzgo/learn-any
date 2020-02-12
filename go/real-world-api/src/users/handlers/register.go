package handlers

import (
	"net/http"

	"real-world-api/src/common"
	UserModel "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

// RegisterForm godoc
type RegisterForm struct {
	User struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
	} `form:"user"`
}

// Register godoc
// @Accept  json
// @Produce  json
// @Router /users/register [post]
func Register(c *gin.Context) {
	var registerForm RegisterForm
	c.ShouldBindJSON(&registerForm)
	// TODO: 查看是否已经存在用户
	userModel := UserModel.New()

	userModel.Email = registerForm.User.Email
	userModel.Username = registerForm.User.Username
	userModel.SetPassword(registerForm.User.Password)

	err := UserModel.CreateUser(userModel)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user = new(UserSchema)
	user.Email = userModel.Email
	user.Username = userModel.Username
	tokenString, err := common.JWTSign()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user.Token = tokenString
	user.Bio = userModel.Bio
	user.Image = userModel.Image
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
	return
}
