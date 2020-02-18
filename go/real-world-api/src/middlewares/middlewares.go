package middlewares

import (
	"log"
	"net/http"
	"real-world-api/src/common"
	"regexp"

	userModels "real-world-api/src/users/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


// HandleServerErrors godoc
func HandleServerErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Internal Error: ", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server Internal Error"})
			}
		}()
		c.Next()
	}
}


// JWTAuth godoc
func JWTAuth() gin.HandlerFunc {
	headerRegex := regexp.MustCompile(`^Token\s*(.*)$`)

	return func (c *gin.Context)  {
		authHeader := c.Request.Header.Get("Authorization")
		matched := headerRegex.FindStringSubmatch(authHeader)
		if (len(matched) == 2) {
			tokenString := matched[1]
			if token, err := common.JWTValidate(tokenString); err == nil {
				claims, ok := token.Claims.(jwt.MapClaims)
				if ok && token.Valid {
					userid := uint(claims["userid"].(float64))
					if user, err := userModels.GetUserByID(userid); *user != (userModels.UserModel{}) && err == nil {
						c.Set(common.KeyJwtCurentUser, user)
						c.Next()
						return
					}
				}
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, common.GenErrorJSON(common.ErrUnauthorized))
		
	}
}

