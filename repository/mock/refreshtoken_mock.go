package mock

import (
	"math/rand"
	"time"
	"webapp-demo/entity"
)

type MockRefreshTokenRepository struct {
	tokens map[entity.EntityId]*entity.RefreshToken
}

func NewMockRefreshTokenRepository() *MockRefreshTokenRepository {
	return &MockRefreshTokenRepository{
		tokens: map[entity.EntityId]*entity.RefreshToken{},
	}
}

func (r *MockRefreshTokenRepository) Create(token string, expiredAt time.Time,
	userId entity.EntityId) (*entity.RefreshToken, error) {

	tokenObj := &entity.RefreshToken{
		TokenId:   entity.EntityId(rand.Int31n(999999)),
		Token:     token,
		ExpiredAt: expiredAt,
		IsActive:  true,
	}
	r.tokens[userId] = tokenObj
	return tokenObj, nil
}

func (r *MockRefreshTokenRepository) Update(token string, expiredAt time.Time,
	userId entity.EntityId) (*entity.RefreshToken, error) {

	tokenObj := r.tokens[userId]
	tokenObj.Token = token
	tokenObj.ExpiredAt = expiredAt
	return tokenObj, nil
}

func (r *MockRefreshTokenRepository) DeactivateByUserId(userId entity.EntityId) error {
	r.tokens[userId].IsActive = false
	return nil
}
