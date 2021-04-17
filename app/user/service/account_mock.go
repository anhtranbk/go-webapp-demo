package service

import (
	"time"
	"webapp-demo/app/user/dto"
	"webapp-demo/app/user/entity"
)

type MockAccountService struct {
}

func NewMockAuthService() *MockAccountService {
	return &MockAccountService{}
}

func (s *MockAccountService) UserSignUp(userReg dto.SignUpDto) (*dto.UserDto, error) {
	return &dto.UserDto{
		UserId:    184615,
		UserName:  "tjeubaoit",
		Email:     "sieunhancamlon@gmail.com",
		Dob:       "2008-10-15",
		Gender:    entity.Male,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (s *MockAccountService) UserSignIn(userLogin dto.SignInDto) (*dto.AccessTokenDto, error) {
	return &dto.AccessTokenDto{
		AccessToken:  "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1ODUwNjMyMTYsIm5iZiI6MTU4NTA2MzIxNiwianRpIj",
		TokenType:    "Bearer",
		ExpiredAt:    time.Now().Add(time.Minute * 60),
		RefreshToken: "7c271bec2288b77c",
	}, nil
}
