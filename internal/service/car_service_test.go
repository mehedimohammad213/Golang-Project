package service

import (
	"errors"
	"testing"

	"github.com/user/car-project/internal/models"
)

// MockRepository is a simple mock for testing
type MockRepository struct {
	cars []models.Car
	err  error
}

func (m *MockRepository) Create(car *models.Car) error  { return m.err }
func (m *MockRepository) GetAll() ([]models.Car, error) { return m.cars, m.err }
func (m *MockRepository) GetByID(id int64) (*models.Car, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &models.Car{ID: id}, nil
}
func (m *MockRepository) Update(car *models.Car) error { return m.err }
func (m *MockRepository) Delete(id int64) error        { return m.err }

func TestGetCarByID(t *testing.T) {
	mockRepo := &MockRepository{}
	svc := NewCarService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		car, err := svc.GetCarByID(1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if car.ID != 1 {
			t.Errorf("Expected ID 1, got %v", car.ID)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		mockRepo.err = errors.New("not found")
		_, err := svc.GetCarByID(1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
