package handler

import (
	"net/http"
	"webapp-demo/dtos"
	"webapp-demo/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service service.AuthService
}

func (h AuthHandler) SignUp(c *fiber.Ctx) error {
	var signUp dtos.SignUpDto
	if err := c.BodyParser(&signUp); err != nil {
		c.Status(http.StatusBadRequest).JSON(dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserRegistration(signUp)
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

	resp, err := h.Service.UserLogin(signIn)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(dtos.E{Error: err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(&resp)
	return nil
}
