package service

import (
	"github.com/user/car-project/internal/models"
	"github.com/user/car-project/internal/repository"
	"github.com/user/car-project/internal/utils"
)

type CarService interface {
	CreateCar(car *models.Car) error
	GetCars() ([]models.Car, error)
	GetCarByID(id int64) (*models.Car, error)
	UpdateCar(car *models.Car) error
	DeleteCar(id int64) error
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}

func (s *carService) CreateCar(car *models.Car) error {
	// Add business logic/validation here if needed
	return s.repo.Create(car)
}

func (s *carService) GetCars() ([]models.Car, error) {
	return s.repo.GetAll()
}

func (s *carService) GetCarByID(id int64) (*models.Car, error) {
	car, err := s.repo.GetByID(id)
	if err != nil {
		return nil, utils.ErrNotFound
	}
	return car, nil
}

func (s *carService) UpdateCar(car *models.Car) error {
	return s.repo.Update(car)
}

func (s *carService) DeleteCar(id int64) error {
	return s.repo.Delete(id)
}
