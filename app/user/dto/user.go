package dto

import (
	"time"
	"webapp-demo/app/user/entity"
)

type UserDto struct {
	UserId    entity.ID     `json:"user_id"`
	UserName  string        `json:"username"`
	Email     string        `json:"email"`
	Dob       string        `json:"dob"`
	Gender    entity.Gender `json:"gender"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
