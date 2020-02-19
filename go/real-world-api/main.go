package main

import (
	"log"
	"real-world-api/src/apis"
	"real-world-api/src/common"
	"real-world-api/src/db"
	"real-world-api/src/middlewares"
	userModels "real-world-api/src/users/models"

	"github.com/gin-contrib/cors"

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

	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4100"}

	r.Use(gin.Logger(), middlewares.HandleServerErrors(), cors.New(config))

	api := r.Group("/api")

	db := db.GetDB()
	db.AutoMigrate(&userModels.UserModel{}, &userModels.FollowModel{})
	defer db.Close()

	// Configure Endpoints
	apis.UseUsersEndpoints(api)

	// Health check
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Swagger Configuration
	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run("localhost:3000")
}
