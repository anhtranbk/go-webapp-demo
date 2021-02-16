package services

import (
	"time"
	"webapp-demo/dtos"
)

func UserRegistration(userReg dtos.UserRegistrationDto) (dtos.UserDto, error) {
	return dtos.UserDto{
		UserId:    184615,
		UserName:  "tjeubaoit",
		Email:     "sieunhancamlon@gmail.com",
		Dob:       "2008-10-15",
		Gender:    dtos.Male,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func UserLogin(userLogin dtos.UserLoginDto) (dtos.UserTokenDto, error) {
	return dtos.UserTokenDto{
		AccessToken:  "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1ODUwNjMyMTYsIm5iZiI6MTU4NTA2MzIxNiwianRpIj",
		TokenType:    "Bearer",
		ExpiredAt:    time.Now().Add(time.Minute * 60),
		RefreshToken: "7c271bec2288b77c",
	}, nil
}
