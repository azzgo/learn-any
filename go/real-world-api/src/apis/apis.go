package apis

import (
	articleHandlers "real-world-api/src/articles/handlers"
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
	api.POST("/profiles/:username/follow", authRequireLoginMiddleware, userHandlers.Follow)
	api.DELETE("/profiles/:username/follow", authRequireLoginMiddleware, userHandlers.UnFollow)
	api.GET("/profiles/:username", authNoNeedLoginMiddleware, userHandlers.Profile)

	// article
	api.GET("/articles", authNoNeedLoginMiddleware, articleHandlers.GetArictles)
	api.GET("/articles/feed", authRequireLoginMiddleware, articleHandlers.GetArticlesByFeed)
}
