package models

import "time"

type CarGrade struct {
	ID            int64     `db:"id" json:"id"`
	CarID         int64     `db:"car_id" json:"car_id"`
	GradeOverall  *string   `db:"grade_overall" json:"grade_overall"`
	GradeExterior *string   `db:"grade_exterior" json:"grade_exterior"`
	GradeInterior *string   `db:"grade_interior" json:"grade_interior"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
