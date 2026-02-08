package models

import "time"

type CarDetail struct {
	ID          int64     `db:"id" json:"id"`
	CarID       int64     `db:"car_id" json:"car_id"`
	ShortTitle  *string   `db:"short_title" json:"short_title"`
	FullTitle   *string   `db:"full_title" json:"full_title"`
	Description *string   `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type CarSubDetail struct {
	ID          int64     `db:"id" json:"id"`
	CarDetailID int64     `db:"car_detail_id" json:"car_detail_id"`
	Title       *string   `db:"title" json:"title"`
	Description *string   `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
