package main

import (
	"log"
	"real-world-api/src/common"
	"real-world-api/src/db"
	"real-world-api/src/users"
	userModels "real-world-api/src/users/models"

	_ "real-world-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/go-sql-driver/mysql"
)

// @title Swagger Example API
// @version 1.0
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	common.InitConfig()

	r := gin.Default()

	api := r.Group("/api")

	db := db.GetDB()
	db.AutoMigrate(&userModels.UserModel{})
	defer db.Close()

	users.UseUsersEndpoints(api.Group("/users"))

	// Health check
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Swagger Configuration
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()
	log.Println("listen and serve on http://127.0.0.1:8080")
}
