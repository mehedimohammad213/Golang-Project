package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/user/car-project/internal/models"
)

type PermissionRepository interface {
	Create(ctx context.Context, permission *models.Permission) error
	GetByID(ctx context.Context, id int64) (*models.Permission, error)
	GetAll(ctx context.Context) ([]models.Permission, error)
	Update(ctx context.Context, permission *models.Permission) error
	Delete(ctx context.Context, id int64) error
	AssignPermissionToRole(ctx context.Context, roleID, permissionID int64) error
	GetByUserID(ctx context.Context, userID int64) ([]string, error)
}

type permissionRepository struct {
	DB *sqlx.DB
}

func NewPermissionRepository(db *sqlx.DB) PermissionRepository {
	return &permissionRepository{DB: db}
}

func (r *permissionRepository) Create(ctx context.Context, permission *models.Permission) error {
	query := `INSERT INTO permissions (name, slug, module, created_at, updated_at) 
			  VALUES (:name, :slug, :module, :created_at, :updated_at) 
			  RETURNING id`

	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRowxContext(ctx, permission).Scan(&id)
	if err != nil {
		return err
	}
	permission.ID = id
	return nil
}

func (r *permissionRepository) GetByID(ctx context.Context, id int64) (*models.Permission, error) {
	var permission models.Permission
	err := r.DB.GetContext(ctx, &permission, "SELECT * FROM permissions WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) GetAll(ctx context.Context) ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.DB.SelectContext(ctx, &permissions, "SELECT * FROM permissions ORDER BY module, name ASC")
	return permissions, err
}

func (r *permissionRepository) Update(ctx context.Context, permission *models.Permission) error {
	query := `UPDATE permissions SET name=:name, module=:module, updated_at=CURRENT_TIMESTAMP WHERE id=:id`
	_, err := r.DB.NamedExecContext(ctx, query, permission)
	return err
}

func (r *permissionRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM permissions WHERE id = $1", id)
	return err
}

func (r *permissionRepository) AssignPermissionToRole(ctx context.Context, roleID, permissionID int64) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO permission_role (role_id, permission_id) VALUES ($1, $2) ON CONFLICT DO NOTHING", roleID, permissionID)
	return err
}

func (r *permissionRepository) GetByUserID(ctx context.Context, userID int64) ([]string, error) {
	query := `
		SELECT DISTINCT p.slug
		FROM permissions p
		JOIN permission_role pr ON p.id = pr.permission_id
		JOIN role_user ru ON pr.role_id = ru.role_id
		WHERE ru.user_id = $1
	`
	var slugs []string
	err := r.DB.SelectContext(ctx, &slugs, query, userID)
	return slugs, err
}
