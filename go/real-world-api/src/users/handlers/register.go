package handlers

import (
	"net/http"

	"real-world-api/src/common"
	"real-world-api/src/errorcodes"
	UserModel "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

// RegisterForm godoc
type RegisterForm struct {
	User struct {
		Username string `form:"username" binding:"required"`
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required,gte=8"`
	} `form:"user"`
}

// Register godoc
// @Accept  json
// @Produce  json
// @Router /users/register [post]
func Register(c *gin.Context) {
	var registerForm RegisterForm
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON("Email, UserName, Password are required"))
		return
	}

	if user, _ := UserModel.GetUserByEmail(registerForm.User.Email); user != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON(errorcodes.UserAlreadyExists))
		return
	}

	userModel, err := UserModel.CreateUser(
		registerForm.User.Email,
		registerForm.User.Username,
		registerForm.User.Password,
	)

	if err != nil {
		panic(err)
	}

	genRegisterResponse(c, userModel)
}

func genRegisterResponse(c *gin.Context, userModel *UserModel.UserModel) {
	var user = new(UserSchema)
	user.Email = userModel.Email
	user.Username = userModel.Username
	user.Bio = userModel.Bio
	user.Image = userModel.Image
	tokenString, err := common.JWTSign()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
	} else {
		user.Token = tokenString
		c.JSON(http.StatusCreated, gin.H{
			"user": user,
		})
	}
}
