package service

import (
	"context"
	"time"

	"github.com/user/car-project/internal/dto"
	"github.com/user/car-project/internal/models"
	"github.com/user/car-project/internal/repository"
	"github.com/user/car-project/internal/utils"
)

type PermissionService interface {
	CreatePermission(ctx context.Context, req dto.CreatePermissionRequest) (*dto.PermissionResponse, error)
	GetPermissions(ctx context.Context) ([]dto.PermissionResponse, error)
	GetPermissionByID(ctx context.Context, id int64) (*dto.PermissionResponse, error)
	UpdatePermission(ctx context.Context, id int64, req dto.UpdatePermissionRequest) error
	DeletePermission(ctx context.Context, id int64) error
	GetUserPermissions(ctx context.Context, userID int64) ([]string, error)
}

type permissionService struct {
	repo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) PermissionService {
	return &permissionService{repo: repo}
}

func (s *permissionService) CreatePermission(ctx context.Context, req dto.CreatePermissionRequest) (*dto.PermissionResponse, error) {
	perm := &models.Permission{
		Name:      req.Name,
		Slug:      req.Slug,
		Module:    req.Module,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, perm); err != nil {
		return nil, err
	}

	return &dto.PermissionResponse{
		ID:        perm.ID,
		Name:      perm.Name,
		Slug:      perm.Slug,
		Module:    perm.Module,
		CreatedAt: perm.CreatedAt,
		UpdatedAt: perm.UpdatedAt,
	}, nil
}

func (s *permissionService) GetPermissions(ctx context.Context) ([]dto.PermissionResponse, error) {
	perms, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var permDTOs []dto.PermissionResponse
	for _, p := range perms {
		permDTOs = append(permDTOs, dto.PermissionResponse{
			ID:        p.ID,
			Name:      p.Name,
			Slug:      p.Slug,
			Module:    p.Module,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return permDTOs, nil
}

func (s *permissionService) GetPermissionByID(ctx context.Context, id int64) (*dto.PermissionResponse, error) {
	perm, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if perm == nil {
		return nil, utils.ErrNotFound
	}

	return &dto.PermissionResponse{
		ID:        perm.ID,
		Name:      perm.Name,
		Slug:      perm.Slug,
		Module:    perm.Module,
		CreatedAt: perm.CreatedAt,
		UpdatedAt: perm.UpdatedAt,
	}, nil
}

func (s *permissionService) UpdatePermission(ctx context.Context, id int64, req dto.UpdatePermissionRequest) error {
	perm, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if perm == nil {
		return utils.ErrNotFound
	}

	if req.Name != nil {
		perm.Name = *req.Name
	}
	if req.Module != nil {
		perm.Module = req.Module
	}

	return s.repo.Update(ctx, perm)
}

func (s *permissionService) DeletePermission(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *permissionService) GetUserPermissions(ctx context.Context, userID int64) ([]string, error) {
	return s.repo.GetByUserID(ctx, userID)
}
