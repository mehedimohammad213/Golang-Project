package models

import "time"

type CarModel struct {
	ID        int64     `db:"id" json:"id"`
	MakeID    int64     `db:"make_id" json:"make_id"`
	Name      string    `db:"name" json:"name"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
