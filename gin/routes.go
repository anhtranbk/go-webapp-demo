package gin

import (
	"webapp-demo/gin/handler"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(router *gin.RouterGroup, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	auth.POST("/registration", authHandler.UserRegistration)
	auth.POST("/login", authHandler.UserLogin)
}
