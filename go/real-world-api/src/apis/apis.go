package apis

import (
	"real-world-api/src/middlewares"
	userHandlers "real-world-api/src/users/handlers"

	"github.com/gin-gonic/gin"
)

// UseUsersEndpoints 拦截 /api/users 的请求
func UseUsersEndpoints(api *gin.RouterGroup) {
	authMiddleware := middlewares.JWTAuth()
	// users
	api.POST("/users/login", userHandlers.Login)
	api.POST("/users/", userHandlers.Register)

	// user
	api.GET("/user",authMiddleware, userHandlers.CurrentUser)
	api.GET("/user/:username", userHandlers.Profile)
}
