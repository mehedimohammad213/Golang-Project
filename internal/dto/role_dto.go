package dto

import "time"

type CreateRoleRequest struct {
	Name        string  `json:"name" binding:"required,min=2,max=100"`
	Slug        string  `json:"slug" binding:"required,min=2,max=100,alphanum"`
	Description *string `json:"description,omitempty"`
}

type UpdateRoleRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=2,max=100"`
	Description *string `json:"description,omitempty"`
}

type RoleResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AssignRoleRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
	RoleID int64 `json:"role_id" binding:"required"`
}
