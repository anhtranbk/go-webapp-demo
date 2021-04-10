package handler

import (
	"net/http"
	"webapp-demo/core"
	"webapp-demo/dto"
	"webapp-demo/service"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AppContext *core.AppContext
	Service    service.AccountService
}

func NewAccountHandler(appCtx *core.AppContext) *AccountHandler {
	return &AccountHandler{
		AppContext: appCtx,
		Service:    service.NewAccountService(appCtx),
	}
}

func (h AccountHandler) SignUp(c *fiber.Ctx) error {
	var signUp dto.SignUpDto
	if err := c.BodyParser(&signUp); err != nil {
		c.Status(http.StatusBadRequest).JSON(dto.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserSignUp(signUp)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(dto.E{Error: err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(&resp)
	return nil
}

func (h AccountHandler) SignIn(c *fiber.Ctx) error {
	var signIn dto.SignInDto
	if err := c.BodyParser(&signIn); err != nil {
		c.Status(http.StatusBadRequest).JSON(dto.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserSignIn(signIn)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(dto.E{Error: err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(&resp)
	return nil
}
