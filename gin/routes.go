package gin

import (
	"webapp-demo/gin/handler"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(router *gin.RouterGroup, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	auth.POST("/signup", authHandler.SignUp)
	auth.POST("/signin", authHandler.SignIn)
}
