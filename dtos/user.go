package dtos

import "time"

type Gender string

const (
	Female Gender = "female"
	Male          = "male"
)

type UserDto struct {
	UserId    uint64    `json:"user_id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Dob       string    `json:"dob"`
	Gender    Gender    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
