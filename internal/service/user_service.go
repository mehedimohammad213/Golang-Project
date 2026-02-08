package service

import (
	"context"
	"time"

	"github.com/user/car-project/internal/dto"
	"github.com/user/car-project/internal/models"
	"github.com/user/car-project/internal/repository"
	"github.com/user/car-project/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUsers(ctx context.Context) ([]dto.UserResponse, error)
	GetUserByID(ctx context.Context, id int64) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id int64, req dto.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id int64) error
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	// Check if already exists
	existingUser, _ := s.repo.GetByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, utils.ErrAlreadyExists
	}

	existingUser, _ = s.repo.GetByUsername(ctx, req.Username)
	if existingUser != nil {
		return nil, utils.ErrAlreadyExists
	}

	// Hash password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         req.Name,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedBytes),
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) GetUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var userDTOs []dto.UserResponse
	for _, u := range users {
		userDTOs = append(userDTOs, dto.UserResponse{
			ID:          u.ID,
			Name:        u.Name,
			Username:    u.Username,
			Email:       u.Email,
			IsActive:    u.IsActive,
			LastLoginAt: u.LastLoginAt,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		})
	}
	return userDTOs, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int64) (*dto.UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.ErrNotFound
	}

	return &dto.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		IsActive:    user.IsActive,
		LastLoginAt: user.LastLoginAt,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int64, req dto.UpdateUserRequest) error {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return utils.ErrNotFound
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	return s.repo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.ErrUnauthorized // Or Invalid Credentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, utils.ErrUnauthorized
	}

	// In a real app, generate JWT here
	token := "dummy-jwt-token"

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			IsActive: user.IsActive,
		},
	}, nil
}
