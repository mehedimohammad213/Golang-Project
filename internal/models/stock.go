package models

import "time"

type Stock struct {
	ID        int64     `db:"id" json:"id"`
	CarID     int64     `db:"car_id" json:"car_id"`
	Quantity  int       `db:"quantity" json:"quantity"`
	Notes     *string   `db:"notes" json:"notes"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
