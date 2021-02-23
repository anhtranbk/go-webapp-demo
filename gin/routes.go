package gin

import (
	"webapp-demo/gin/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine, authHander *handlers.AuthHandler) {
	InitAuthRoutes(app, authHander)
}

func InitAuthRoutes(app *gin.Engine, handler *handlers.AuthHandler) {
	v1 := app.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/registration", handler.UserRegistration)
	auth.POST("/login", handler.UserLogin)
}
