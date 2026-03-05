package models

import "time"

// RAGChunk is a single indexed document chunk with its embedding.
type RAGChunk struct {
	ID         int64     `db:"id" json:"id"`
	SourceType string    `db:"source_type" json:"source_type"`
	SourceID   string    `db:"source_id" json:"source_id"`
	Content    string    `db:"content" json:"content"`
	Metadata   string    `db:"metadata" json:"metadata"` // JSON string
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
