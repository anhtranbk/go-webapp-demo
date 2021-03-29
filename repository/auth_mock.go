package repository

import (
	"errors"
	"fmt"
	"math/rand"
	"webapp-demo/entity"
)

type MockUserRepository struct {
	users map[entity.EntityId]*entity.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[entity.EntityId]*entity.User),
	}
}

func (r *MockUserRepository) Create(user *entity.User) (*entity.User, error) {
	user.UserId = entity.EntityId(rand.Int31n(999999))
	if user.UserName == "" {
		user.UserName = fmt.Sprintf("user_%d", user.UserId)
	}
	r.users[user.UserId] = user
	return user, nil
}

func (r *MockUserRepository) Update(user *entity.User) error {
	r.users[user.UserId] = user
	return nil
}

func (r *MockUserRepository) Delete(userId entity.EntityId) error {
	delete(r.users, userId)
	return nil
}

func (r *MockUserRepository) FindAll(offset int32, limit int32) []*entity.User {
	values := make([]*entity.User, 0, len(r.users))
	for _, val := range r.users {
		values = append(values, val)
	}
	return values
}

func (r *MockUserRepository) FindById(userId entity.EntityId) (*entity.User, error) {
	return r.users[userId], nil
}

func (r *MockUserRepository) FindByEmail(email string) (*entity.User, error) {
	for _, val := range r.users {
		if val.Email == email {
			return val, nil
		}
	}
	return nil, errors.New("Not found")
}

func (r *MockUserRepository) IsExist(email string) (bool, error) {
	for _, val := range r.users {
		if val.Email == email {
			return true, nil
		}
	}
	return false, nil
}
