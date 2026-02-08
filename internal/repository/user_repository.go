package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/user/car-project/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int64) error
}

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (name, username, email, password_hash, is_active, created_at, updated_at) 
			  VALUES (:name, :username, :email, :password_hash, :is_active, :created_at, :updated_at) 
			  RETURNING id`

	// Use NamedQuery to get back the ID
	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRowxContext(ctx, user).Scan(&id)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := r.DB.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Or custom ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.DB.GetContext(ctx, &user, "SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.DB.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1 AND deleted_at IS NULL", username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.DB.SelectContext(ctx, &users, "SELECT * FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC")
	return users, err
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET name=:name, email=:email, is_active=:is_active, updated_at=CURRENT_TIMESTAMP 
			  WHERE id=:id AND deleted_at IS NULL`

	_, err := r.DB.NamedExecContext(ctx, query, user)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	// Soft delete
	_, err := r.DB.ExecContext(ctx, "UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1", id)
	return err
}
