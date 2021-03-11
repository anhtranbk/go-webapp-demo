package handlers

import (
	"net/http"
	"webapp-demo/dtos"
	"webapp-demo/service"

	echo "github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service service.AuthService
}

func (h AuthHandler) UserRegistration(c echo.Context) error {
	var userReg dtos.UserRegistrationDto
	if err := c.Bind(&userReg); err != nil {
		c.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserRegistration(userReg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, &resp)
	return nil
}

func (h AuthHandler) UserLogin(c echo.Context) error {
	var userLogin dtos.UserLoginDto
	if err := c.Bind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserLogin(userLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, &resp)
	return nil
}
