package main

import (
	"fmt"
	"real-world-api/src/users"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")

	users.UseUsersEndpoints(api.Group("/users"))

	// health check
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
	fmt.Println("listen and serve on http://127.0.0.1:8080")
}
