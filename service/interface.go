package service

import (
	"webapp-demo/dtos"
)

type AuthService interface {
	UserRegistration(userReg dtos.UserRegistrationDto) (dtos.UserDto, error)
	UserLogin(userLogin dtos.UserLoginDto) (dtos.UserTokenDto, error)
}
