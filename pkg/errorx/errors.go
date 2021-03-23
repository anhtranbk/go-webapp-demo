package errorx

import "errors"

var (
	EmailExisted        = errors.New("The email is already in use with another account")
	InvalidCredentials  = errors.New("The email or password is incorrect")
	InvalidRefreshToken = errors.New("The refresh token is incorrect")
)
