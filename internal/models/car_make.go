package models

import "time"

type CarMake struct {
	ID            int64     `db:"id" json:"id"`
	Name          string    `db:"name" json:"name"`
	OriginCountry *string   `db:"origin_country" json:"origin_country"`
	Status        string    `db:"status" json:"status"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
