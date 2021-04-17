package app

import (
	"webapp-demo/app/user/repository"
	"webapp-demo/app/user/repository/mock"
)

type Repositories struct {
	UserRepo         repository.UserRepo
	RefreshTokenRepo repository.RefreshTokenRepo
}

func NewMockRepositories() *Repositories {
	return &Repositories{
		UserRepo:         mock.NewUserRepo(),
		RefreshTokenRepo: mock.NewRefreshTokenRepo(),
	}
}
