package dto

import "time"

type CreatePermissionRequest struct {
	Name   string  `json:"name" binding:"required,min=2,max=150"`
	Slug   string  `json:"slug" binding:"required,min=2,max=150"`
	Module *string `json:"module,omitempty"`
}

type UpdatePermissionRequest struct {
	Name   *string `json:"name,omitempty" binding:"omitempty,min=2,max=150"`
	Module *string `json:"module,omitempty"`
}

type PermissionResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Module    *string   `json:"module,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
