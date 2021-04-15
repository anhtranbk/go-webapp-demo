package errorx

import "errors"

var (
	// signup errors
	EmailExisted      = errors.New("The email is already in use with another account")
	MalformedPassword = errors.New("Password length must be between 6 and 32 characters")
	MalformedEmail    = errors.New("Email is malformed")
	// signin errors
	InvalidCredentials  = errors.New("The email or password is incorrect")
	InvalidRefreshToken = errors.New("The refresh token is incorrect")
)
