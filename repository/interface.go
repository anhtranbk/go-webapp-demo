package repository

import (
	"time"
	"webapp-demo/entity"
)

type UserRepo interface {
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) error
	Delete(userId entity.ID) error

	FindAll(offset int32, limit int32) []*entity.User
	FindById(userId entity.ID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	IsExist(email string) (bool, error)
}

type RefreshTokenRepo interface {
	Create(token string, expiredAt time.Time, userId entity.ID) (*entity.RefreshToken, error)
	Update(token string, expiredAt time.Time, userId entity.ID) (*entity.RefreshToken, error)
	Deactivate(userId entity.ID) error
}

type Repositories struct {
	UserRepo         UserRepo
	RefreshTokenRepo RefreshTokenRepo
}
