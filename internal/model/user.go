package model

import "time"

type UserModel struct {
	Model
	Username  string `json:"username,omitempty" validate:"required,max=50,min=1"`
	Password  string `json:"password,omitempty" validate:"required,max=255,min=6"`
	IsAdmin   bool
	IsActive  bool
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type RegisterUserModel struct {
	UserModel
	FullName string `json:"full_name,omitempty" validate:"required,max=100,min=1"`
	Email    string `json:"email,omitempty" validate:"required,email"`
}
