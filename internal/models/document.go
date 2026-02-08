package models

import "time"

type Document struct {
	ID           int64     `db:"id" json:"id"`
	CarID        int64     `db:"car_id" json:"car_id"`
	DocumentType string    `db:"document_type" json:"document_type"`
	FileName     string    `db:"file_name" json:"file_name"`
	FilePath     string    `db:"file_path" json:"file_path"`
	FileSize     *int64    `db:"file_size" json:"file_size"`
	MimeType     *string   `db:"mime_type" json:"mime_type"`
	IsPrimary    bool      `db:"is_primary" json:"is_primary"`
	IsHidden     bool      `db:"is_hidden" json:"is_hidden"`
	SortOrder    int       `db:"sort_order" json:"sort_order"`
	UploadedBy   *int64    `db:"uploaded_by" json:"uploaded_by"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
