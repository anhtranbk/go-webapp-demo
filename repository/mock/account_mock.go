package mock

import (
	"errors"
	"fmt"
	"math/rand"
	"webapp-demo/entity"
)

type MockUserRepo struct {
	users map[entity.ID]*entity.User
}

func NewUserRepo() *MockUserRepo {
	return &MockUserRepo{
		users: make(map[entity.ID]*entity.User),
	}
}

func (r *MockUserRepo) Create(user *entity.User) (*entity.User, error) {
	user.UserId = entity.ID(rand.Int31n(999999))
	if user.UserName == "" {
		user.UserName = fmt.Sprintf("user_%d", user.UserId)
	}
	r.users[user.UserId] = user
	return user, nil
}

func (r *MockUserRepo) Update(user *entity.User) error {
	r.users[user.UserId] = user
	return nil
}

func (r *MockUserRepo) Delete(userId entity.ID) error {
	delete(r.users, userId)
	return nil
}

func (r *MockUserRepo) FindAll(offset int32, limit int32) []*entity.User {
	values := make([]*entity.User, 0, len(r.users))
	for _, val := range r.users {
		values = append(values, val)
	}
	return values
}

func (r *MockUserRepo) FindById(userId entity.ID) (*entity.User, error) {
	return r.users[userId], nil
}

func (r *MockUserRepo) FindByEmail(email string) (*entity.User, error) {
	for _, val := range r.users {
		if val.Email == email {
			return val, nil
		}
	}
	return nil, errors.New("Not found")
}

func (r *MockUserRepo) IsExist(email string) (bool, error) {
	for _, val := range r.users {
		if val.Email == email {
			return true, nil
		}
	}
	return false, nil
}
