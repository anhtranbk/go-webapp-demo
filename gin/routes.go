package gin

import (
	"webapp-demo/gin/handler"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(app *gin.Engine, authHandler *handler.AuthHandler) {
	v1 := app.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/registration", authHandler.UserRegistration)
	auth.POST("/login", authHandler.UserLogin)
}
