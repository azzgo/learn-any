package handlers

import (
	"net/http"
	"real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	User struct {
		Username string `form: "username"`
		Email    string `form: "email"`
		Password string `form: "password"`
	} `form: "user"`
}

func Register(c *gin.Context) {
	var registerForm RegisterForm
	c.ShouldBindJSON(&registerForm)
	// TODO: 注册逻辑

	var user = new(models.UserModel)
	user.User.Email = "fakeEmail"
	user.User.Token = "fakeToken"
	user.User.Username = "fakeUsername"
	user.User.Bio = "fakeBio"
	user.User.Image = "fakeImage"
	c.JSON(http.StatusCreated, user)

}
