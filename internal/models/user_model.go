package models

import (
	db "gitlab/go-prolog-api/example/db/sqlc"
	"time"
)

type CreateUserRequest struct {
	Fullname  string `form:"fullname" binding:"required,min=3,max=32,alphanum"`
	Photo     string `form:"photo" binding:"required"`
	Email     string `form:"email" binding:"required,email"`
	Phone     string `form:"phone" binding:"required"`
	Role      string `form:"role" binding:"required"`
	Password  string `form:"password" binding:"required,min=8,max=32"`
}

type UserResponse struct {
	Username  string    `json:"username"`
	Positions string    `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func NewUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username:  user.Fullname,
		Positions: user.Role,
		CreatedAt: user.Createdat.Time,
		UpdateAt:  user.Updatedat.Time,
	}
}

type LoginUserRequest struct {
	Email    string `form:"email" binding:"required,min=3,max=32,email"`
	Password string `form:"password" binding:"required,min=8,max=32"`
}

type LoginUserResponse struct {
	AccsesToken string       `json:"access_token"`
	User        UserResponse `json:"data"`
}
