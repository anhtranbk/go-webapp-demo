package fiber

import (
	"webapp-demo/core"
	"webapp-demo/fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(f *fiber.App, appCtx *core.AppContext) {
	v1 := f.Group("/v1")

	initAccountRoutes(v1, appCtx)
}

func initAccountRoutes(router fiber.Router, appCtx *core.AppContext) {
	auth := router.Group("/auth")
	handler := handler.NewAccountHandler(appCtx)

	auth.Post("/signup", handler.SignUp)
	auth.Post("/signin", handler.SignIn)
}
