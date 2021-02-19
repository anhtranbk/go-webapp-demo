package handlers

import (
	"net/http"
	"webapp-demo/dtos"
	"webapp-demo/services"

	echo "github.com/labstack/echo/v4"
)

func UserRegistration(c echo.Context) error {
	var userReg dtos.UserRegistrationDto
	if err := c.Bind(&userReg); err != nil {
		c.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return err
	}

	resp, err := services.UserRegistration(userReg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, &resp)
	return nil
}

func UserLogin(c echo.Context) error {
	var userLogin dtos.UserLoginDto
	if err := c.Bind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return err
	}

	resp, err := services.UserLogin(userLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, &resp)
	return nil
}
