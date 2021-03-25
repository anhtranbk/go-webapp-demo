package repository

import (
	"time"
	"webapp-demo/entity"
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) error
	Delete(userId entity.EntityId) error

	FindAll(offset int32, limit int32) []*entity.User
	FindById(userId entity.EntityId) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	IsExist(email string) (bool, error)
}

type RefreshTokenRepository interface {
	Create(token string, expiredAt time.Time, userId entity.EntityId) (*entity.RefreshToken, error)
	Update(token string, expiredAt time.Time, userId entity.EntityId) (*entity.RefreshToken, error)
	DeactivateByUserId(userId entity.EntityId) error
}

type Repositories struct {
	UserRepo         UserRepository
	RefreshTokenRepo RefreshTokenRepository
}
