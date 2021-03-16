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
	var userReg dtos.UserRegistrationDto
	if err := c.BodyParser(&userReg); err != nil {
		c.Status(http.StatusBadRequest).JSON(dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserRegistration(userReg)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(dtos.E{Error: err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(&resp)
	return nil
}

func (h AuthHandler) SignIn(c *fiber.Ctx) error {
	var userLogin dtos.UserLoginDto
	if err := c.BodyParser(&userLogin); err != nil {
		c.Status(http.StatusBadRequest).JSON(dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserLogin(userLogin)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(dtos.E{Error: err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(&resp)
	return nil
}
