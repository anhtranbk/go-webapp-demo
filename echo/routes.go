package echo

import (
	"webapp-demo/core"
	"webapp-demo/echo/handler"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, appCtx *core.AppContext) {
	v1 := e.Group("/v1")

	initAuthRoutes(v1, appCtx)
}

func initAuthRoutes(router *echo.Group, appCtx *core.AppContext) {
	auth := router.Group("/auth")
	handler := handler.NewAuthHandler(appCtx)

	auth.POST("/signup", handler.SignUp)
	auth.POST("/signin", handler.SignIn)
}
