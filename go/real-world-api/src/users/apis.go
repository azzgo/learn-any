package users

import (
	"github.com/gin-gonic/gin"
	"real-world-api/src/users/handlers"
)

// UseUsersEndpoints 拦截 /api/users 的请求
func UseUsersEndpoints(api *gin.RouterGroup) {
	api.POST("/login", handlers.Login)
	api.POST("/register", handlers.Register)
}
