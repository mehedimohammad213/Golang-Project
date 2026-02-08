package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/user/car-project/internal/models"
)

type CarRepository interface {
	Create(car *models.Car) error
	GetAll() ([]models.Car, error)
	GetByID(id int64) (*models.Car, error)
	Update(car *models.Car) error
	Delete(id int64) error
}

type carRepository struct {
	DB *sqlx.DB
}

func NewCarRepository(db *sqlx.DB) CarRepository {
	return &carRepository{DB: db}
}

func (r *carRepository) Create(car *models.Car) error {
	query := `INSERT INTO cars (model_id, ref_no, package, body_type, year, color, reg_year_month, mileage_km, chassis_no_full, engine_cc, fuel, transmission, drive, engine_number, seats, number_of_keys, keys_feature, steering, location, country_origin, status) 
			  VALUES (:model_id, :ref_no, :package, :body_type, :year, :color, :reg_year_month, :mileage_km, :chassis_no_full, :engine_cc, :fuel, :transmission, :drive, :engine_number, :seats, :number_of_keys, :keys_feature, :steering, :location, :country_origin, :status) 
			  RETURNING id`

	rows, err := r.DB.NamedQuery(query, car)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&car.ID)
	}
	return nil
}

func (r *carRepository) GetAll() ([]models.Car, error) {
	var cars []models.Car
	err := r.DB.Select(&cars, "SELECT * FROM cars ORDER BY created_at DESC")
	return cars, err
}

func (r *carRepository) GetByID(id int64) (*models.Car, error) {
	var car models.Car
	err := r.DB.Get(&car, "SELECT * FROM cars WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (r *carRepository) Update(car *models.Car) error {
	query := `UPDATE cars SET model_id=:model_id, ref_no=:ref_no, package=:package, body_type=:body_type, year=:year, color=:color, 
			  reg_year_month=:reg_year_month, mileage_km=:mileage_km, chassis_no_full=:chassis_no_full, engine_cc=:engine_cc, 
			  fuel=:fuel, transmission=:transmission, drive=:drive, engine_number=:engine_number, seats=:seats, 
			  number_of_keys=:number_of_keys, keys_feature=:keys_feature, steering=:steering, location=:location, 
              country_origin=:country_origin, status=:status, updated_at=CURRENT_TIMESTAMP 
			  WHERE id=:id`

	_, err := r.DB.NamedExec(query, car)
	return err
}

func (r *carRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM cars WHERE id = $1", id)
	return err
}
