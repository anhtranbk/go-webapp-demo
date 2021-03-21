package service

import (
	"webapp-demo/dtos"
)

type AuthService interface {
	UserSignUp(signUp dtos.SignUpDto) (*dtos.UserDto, error)
	UserSignIn(signIn dtos.SignInDto) (*dtos.AccessTokenDto, error)
}
