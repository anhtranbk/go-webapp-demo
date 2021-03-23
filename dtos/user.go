package dtos

import (
	"time"
	"webapp-demo/entity"
)

type UserDto struct {
	UserId    entity.EntityId `json:"user_id"`
	UserName  string          `json:"username"`
	Email     string          `json:"email"`
	Dob       string          `json:"dob"`
	Gender    entity.Gender   `json:"gender"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
