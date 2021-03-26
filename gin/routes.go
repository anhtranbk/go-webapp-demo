package gin

import (
	"webapp-demo/core"
	"webapp-demo/gin/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine, appCtx *core.AppContext) {
	v1 := engine.Group("/v1")

	initAuthRoutes(v1, appCtx)
}

func initAuthRoutes(router *gin.RouterGroup, appCtx *core.AppContext) {
	auth := router.Group("/auth")
	// handler := &handler.AuthHandler{Service: service.NewMockAuthService()}
	handler := handler.NewAuthHandler(appCtx)

	auth.POST("/signup", handler.SignUp)
	auth.POST("/signin", handler.SignIn)
}
