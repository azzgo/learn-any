package main

import (
	"log"
	"real-world-api/src/users"

	_ "real-world-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @BasePath /api
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

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url)) 	

	r.Run()
	log.Println("listen and serve on http://127.0.0.1:8080")
}
