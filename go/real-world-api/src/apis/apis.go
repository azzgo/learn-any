package apis

import (
	UserHandlers "real-world-api/src/users/handlers"

	"github.com/gin-gonic/gin"
)

// UseUsersEndpoints 拦截 /api/users 的请求
func UseUsersEndpoints(api *gin.RouterGroup) {
	api.POST("/users/login", UserHandlers.Login)
	api.POST("/users/", UserHandlers.Register)
}
