package models

import (
	"time"
)

type Cart struct {
	ID        int64     `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	CarID     int64     `db:"car_id" json:"car_id"`
	Quantity  int       `db:"quantity" json:"quantity"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Order struct {
	ID              int64     `db:"id" json:"id"`
	UserID          int64     `db:"user_id" json:"user_id"`
	TotalAmount     float64   `db:"total_amount" json:"total_amount"`
	ShippingAddress *string   `db:"shipping_address" json:"shipping_address"`
	Status          string    `db:"status" json:"status"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type OrderItem struct {
	ID        int64     `db:"id" json:"id"`
	OrderID   int64     `db:"order_id" json:"order_id"`
	CarID     int64     `db:"car_id" json:"car_id"`
	Quantity  int       `db:"quantity" json:"quantity"`
	Price     float64   `db:"price" json:"price"`
	Notes     *string   `db:"notes" json:"notes"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
