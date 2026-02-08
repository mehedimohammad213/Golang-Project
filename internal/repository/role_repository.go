package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/user/car-project/internal/models"
)

type RoleRepository interface {
	Create(ctx context.Context, role *models.Role) error
	GetByID(ctx context.Context, id int64) (*models.Role, error)
	GetAll(ctx context.Context) ([]models.Role, error)
	Update(ctx context.Context, role *models.Role) error
	Delete(ctx context.Context, id int64) error
	AssignRoleToUser(ctx context.Context, userID, roleID int64) error
}

type roleRepository struct {
	DB *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) RoleRepository {
	return &roleRepository{DB: db}
}

func (r *roleRepository) Create(ctx context.Context, role *models.Role) error {
	query := `INSERT INTO roles (name, slug, description, created_at, updated_at) 
			  VALUES (:name, :slug, :description, :created_at, :updated_at) 
			  RETURNING id`

	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRowxContext(ctx, role).Scan(&id)
	if err != nil {
		return err
	}
	role.ID = id
	return nil
}

func (r *roleRepository) GetByID(ctx context.Context, id int64) (*models.Role, error) {
	var role models.Role
	err := r.DB.GetContext(ctx, &role, "SELECT * FROM roles WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetAll(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.SelectContext(ctx, &roles, "SELECT * FROM roles ORDER BY name ASC")
	return roles, err
}

func (r *roleRepository) Update(ctx context.Context, role *models.Role) error {
	query := `UPDATE roles SET name=:name, description=:description, updated_at=CURRENT_TIMESTAMP WHERE id=:id`
	_, err := r.DB.NamedExecContext(ctx, query, role)
	return err
}

func (r *roleRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM roles WHERE id = $1", id)
	return err
}

func (r *roleRepository) AssignRoleToUser(ctx context.Context, userID, roleID int64) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO role_user (user_id, role_id) VALUES ($1, $2) ON CONFLICT DO NOTHING", userID, roleID)
	return err
}
