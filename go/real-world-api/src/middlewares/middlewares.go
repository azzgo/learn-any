package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

// HandleSeverErrors godoc
func HandleSeverErrors() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Next()

		privateErrs := c.Errors.ByType(gin.ErrorTypePrivate)
		if len(privateErrs) != 0 {
			log.Fatal(privateErrs.Last().Error())
			c.JSON(500, gin.H{
				"error": "server internal serror",
			})
		}
	}
}