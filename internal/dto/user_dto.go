package dto

import "time"

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=150"`
	Username string `json:"username" binding:"required,min=3,max=80,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name,omitempty" binding:"omitempty,min=2,max=150"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	IsActive *bool   `json:"is_active,omitempty"`
}

type UserResponse struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	IsActive    bool       `json:"is_active"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
