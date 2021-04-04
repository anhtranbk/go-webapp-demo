package handler

import (
	"net/http"
	"webapp-demo/core"
	"webapp-demo/dtos"
	"webapp-demo/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AppContext *core.AppContext
	Service    service.AuthService
}

func NewAuthHandler(appCtx *core.AppContext) *AuthHandler {
	return &AuthHandler{
		AppContext: appCtx,
		Service:    service.NewAuthService(appCtx),
	}
}

func (h AuthHandler) SignUp(c *fiber.Ctx) error {
	var signUp dtos.SignUpDto
	if err := c.BodyParser(&signUp); err != nil {
		c.Status(http.StatusBadRequest).JSON(dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserSignUp(signUp)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(dtos.E{Error: err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(&resp)
	return nil
}

func (h AuthHandler) SignIn(c *fiber.Ctx) error {
	var signIn dtos.SignInDto
	if err := c.BodyParser(&signIn); err != nil {
		c.Status(http.StatusBadRequest).JSON(dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserSignIn(signIn)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(dtos.E{Error: err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(&resp)
	return nil
}
