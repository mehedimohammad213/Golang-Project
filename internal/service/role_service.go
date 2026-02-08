package service

import (
	"context"
	"time"

	"github.com/user/car-project/internal/dto"
	"github.com/user/car-project/internal/models"
	"github.com/user/car-project/internal/repository"
	"github.com/user/car-project/internal/utils"
)

type RoleService interface {
	CreateRole(ctx context.Context, req dto.CreateRoleRequest) (*dto.RoleResponse, error)
	GetRoles(ctx context.Context) ([]dto.RoleResponse, error)
	GetRoleByID(ctx context.Context, id int64) (*dto.RoleResponse, error)
	UpdateRole(ctx context.Context, id int64, req dto.UpdateRoleRequest) error
	DeleteRole(ctx context.Context, id int64) error
	AssignRole(ctx context.Context, req dto.AssignRoleRequest) error
}

type roleService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) CreateRole(ctx context.Context, req dto.CreateRoleRequest) (*dto.RoleResponse, error) {
	role := &models.Role{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, role); err != nil {
		return nil, err
	}

	return &dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Slug:        role.Slug,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

func (s *roleService) GetRoles(ctx context.Context) ([]dto.RoleResponse, error) {
	roles, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var roleDTOs []dto.RoleResponse
	for _, r := range roles {
		roleDTOs = append(roleDTOs, dto.RoleResponse{
			ID:          r.ID,
			Name:        r.Name,
			Slug:        r.Slug,
			Description: r.Description,
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		})
	}
	return roleDTOs, nil
}

func (s *roleService) GetRoleByID(ctx context.Context, id int64) (*dto.RoleResponse, error) {
	role, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, utils.ErrNotFound
	}

	return &dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Slug:        role.Slug,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

func (s *roleService) UpdateRole(ctx context.Context, id int64, req dto.UpdateRoleRequest) error {
	role, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if role == nil {
		return utils.ErrNotFound
	}

	if req.Name != nil {
		role.Name = *req.Name
	}
	if req.Description != nil {
		role.Description = req.Description
	}

	return s.repo.Update(ctx, role)
}

func (s *roleService) DeleteRole(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *roleService) AssignRole(ctx context.Context, req dto.AssignRoleRequest) error {
	return s.repo.AssignRoleToUser(ctx, req.UserID, req.RoleID)
}
