package middlewares

import (
	"log"
	"net/http"

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
