package models

import "time"

type Car struct {
	ID            int64     `db:"id" json:"id"`
	ModelID       int64     `db:"model_id" json:"model_id" binding:"required"`
	RefNo         *string   `db:"ref_no" json:"ref_no" binding:"required"`
	Package       *string   `db:"package" json:"package"`
	BodyType      *string   `db:"body_type" json:"body_type"`
	Year          *int16    `db:"year" json:"year"`
	Color         *string   `db:"color" json:"color"`
	RegYearMonth  *string   `db:"reg_year_month" json:"reg_year_month"`
	MileageKM     *int32    `db:"mileage_km" json:"mileage_km"`
	ChassisNoFull *string   `db:"chassis_no_full" json:"chassis_no_full"`
	EngineCC      *int32    `db:"engine_cc" json:"engine_cc"`
	Fuel          *string   `db:"fuel" json:"fuel"`
	Transmission  *string   `db:"transmission" json:"transmission"`
	Drive         *string   `db:"drive" json:"drive"`
	EngineNumber  *string   `db:"engine_number" json:"engine_number"`
	Seats         *int16    `db:"seats" json:"seats"`
	NumberOfKeys  *int32    `db:"number_of_keys" json:"number_of_keys"`
	KeysFeature   *string   `db:"keys_feature" json:"keys_feature"`
	Steering      *string   `db:"steering" json:"steering"`
	Location      *string   `db:"location" json:"location"`
	CountryOrigin *string   `db:"country_origin" json:"country_origin"`
	Status        *string   `db:"status" json:"status" binding:"omitempty,oneof=available sold pending"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
