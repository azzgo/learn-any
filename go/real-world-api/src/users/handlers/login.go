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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	user, err := UserModel.GetUserByEmail(form.User.Email)

	if err != nil {
		if err.Error() == "record not found" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "user not exists",
			})
			return
		}

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if common.Hash(form.User.Password) != user.Password {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var resUser = new(UserSchema)
	resUser.Email = user.Email
	resUser.Username = user.Username
	tokenString, err := common.JWTSign()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	resUser.Token = tokenString
	resUser.Bio = user.Bio
	resUser.Image = user.Image

	c.JSON(http.StatusOK, gin.H{"user": resUser})
	return
}
