package apis

import (
	"real-world-api/src/middlewares"
	userHandlers "real-world-api/src/users/handlers"

	"github.com/gin-gonic/gin"
)

// UseUsersEndpoints 拦截 /api/users 的请求
func UseUsersEndpoints(api *gin.RouterGroup) {
	authRequireLoginMiddleware := middlewares.JWTAuth(true)
	authNoNeedLoginMiddleware := middlewares.JWTAuth(false)

	// users
	api.POST("/users/login", userHandlers.Login)
	api.POST("/users/", userHandlers.Register)

	// user
	api.GET("/user",authRequireLoginMiddleware, userHandlers.CurrentUser)
	api.GET("/profiles/:username", authNoNeedLoginMiddleware, userHandlers.Profile)
}
