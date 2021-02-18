package handlers

import (
	"net/http"
	"webapp-demo/dtos"
	"webapp-demo/services"

	"github.com/gofiber/fiber/v2"
)

func UserRegistration(c *fiber.Ctx) error {
	var userReg dtos.UserRegistrationDto
	if err := c.BodyParser(&userReg); err != nil {
		c.Status(http.StatusBadRequest).JSON(dtos.E{Error: err.Error()})
		return err
	}

	resp, err := services.UserRegistration(userReg)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(dtos.E{Error: err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(&resp)
	return nil
}

func UserLogin(c *fiber.Ctx) error {
	var userLogin dtos.UserLoginDto
	if err := c.BodyParser(&userLogin); err != nil {
		c.Status(http.StatusBadRequest).JSON(dtos.E{Error: err.Error()})
		return err
	}

	resp, err := services.UserLogin(userLogin)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(dtos.E{Error: err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(&resp)
	return nil
}
