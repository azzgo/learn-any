package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User struct {
		Email    string `form: 'email'`
		Password string `form: 'password'`
	} `form: 'user'`
}

// Login godoc
// @Accept  json
// @Produce  json
// @Produce  json
// @Router /users/login [post]
func Login(c *gin.Context) {
	var form LoginForm
	if c.ShouldBindJSON(&form) == nil {
		log.Println(form.User.Email)
		log.Println(form.User.Password)
	}

	// 验证逻辑
	if form.User.Email == "test@gl.com" && form.User.Password == "123" {
		c.JSON(http.StatusOK, gin.H{"user": map[string]string{
			"Email":    "fake@email.com",
			"Username": "fakeUser",
			"Bio":      "bio",
			"Image":    "Image",
		}})
		return
	}

	c.String(http.StatusUnauthorized, "Unauthorized")
}
