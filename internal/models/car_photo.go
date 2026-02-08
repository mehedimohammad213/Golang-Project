package models

import "time"

type CarPhoto struct {
	ID        int64     `db:"id" json:"id"`
	CarID     int64     `db:"car_id" json:"car_id"`
	URL       string    `db:"url" json:"url"`
	IsPrimary bool      `db:"is_primary" json:"is_primary"`
	SortOrder int       `db:"sort_order" json:"sort_order"`
	IsHidden  bool      `db:"is_hidden" json:"is_hidden"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
