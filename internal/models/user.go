package models

import "time"

type User struct {
	ID           int64      `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Username     string     `db:"username" json:"username"`
	Email        string     `db:"email" json:"email"`
	PasswordHash string     `db:"password_hash" json:"-"`
	IsActive     bool       `db:"is_active" json:"is_active"`
	LastLoginAt  *time.Time `db:"last_login_at" json:"last_login_at"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	CreatedBy    *int64     `db:"created_by" json:"created_by"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	UpdatedBy    *int64     `db:"updated_by" json:"updated_by"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at"`
}

type RoleUser struct {
	ID         int64     `db:"id" json:"id"`
	UserID     int64     `db:"user_id" json:"user_id"`
	RoleID     int64     `db:"role_id" json:"role_id"`
	AssignedAt time.Time `db:"assigned_at" json:"assigned_at"`
	AssignedBy *int64    `db:"assigned_by" json:"assigned_by"`
}
