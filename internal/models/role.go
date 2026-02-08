package models

import "time"

type Role struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Slug        string    `db:"slug" json:"slug"`
	Description *string   `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Permission struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Slug      string    `db:"slug" json:"slug"`
	Module    *string   `db:"module" json:"module"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type PermissionRole struct {
	ID           int64     `db:"id" json:"id"`
	PermissionID int64     `db:"permission_id" json:"permission_id"`
	RoleID       int64     `db:"role_id" json:"role_id"`
	AssignedAt   time.Time `db:"assigned_at" json:"assigned_at"`
}
