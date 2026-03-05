package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/user/car-project/internal/models"
)

// RAGChunkInput is used for inserting chunks (embedding passed as pgvector string).
type RAGChunkInput struct {
	SourceType string
	SourceID   string
	Content    string
	Embedding  string // pgvector format "[f1,f2,...]"
	Metadata   string // JSON object string, e.g. "{}"
}

type RAGRepository interface {
	UpsertChunks(ctx context.Context, chunks []RAGChunkInput) error
	SearchByEmbedding(ctx context.Context, embedding string, topK int) ([]models.RAGChunk, error)
	DeleteBySource(ctx context.Context, sourceType, sourceID string) error
	DeleteAllBySourceType(ctx context.Context, sourceType string) error
	GetCarsContentForIndexing(ctx context.Context) ([]CarContentRow, error)
}

// CarContentRow holds one car's aggregated text for RAG indexing.
type CarContentRow struct {
	CarID   int64  `db:"car_id"`
	Content string `db:"content"`
}

type ragRepository struct {
	DB *sqlx.DB
}

func NewRAGRepository(db *sqlx.DB) RAGRepository {
	return &ragRepository{DB: db}
}

func (r *ragRepository) UpsertChunks(ctx context.Context, chunks []RAGChunkInput) error {
	for _, c := range chunks {
		_, err := r.DB.ExecContext(ctx,
			`INSERT INTO rag_chunks (source_type, source_id, content, embedding, metadata)
			 VALUES ($1, $2, $3, $4::vector, $5::jsonb)`,
			c.SourceType, c.SourceID, c.Content, c.Embedding, c.Metadata,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ragRepository) SearchByEmbedding(ctx context.Context, embedding string, topK int) ([]models.RAGChunk, error) {
	if topK <= 0 {
		topK = 5
	}
	var chunks []models.RAGChunk
	err := r.DB.SelectContext(ctx, &chunks,
		`SELECT id, source_type, source_id, content, COALESCE(metadata::text, '{}') AS metadata, created_at
		 FROM rag_chunks
		 WHERE embedding IS NOT NULL
		 ORDER BY embedding <=> $1::vector
		 LIMIT $2`,
		embedding, topK,
	)
	return chunks, err
}

func (r *ragRepository) DeleteBySource(ctx context.Context, sourceType, sourceID string) error {
	_, err := r.DB.ExecContext(ctx,
		`DELETE FROM rag_chunks WHERE source_type = $1 AND source_id = $2`,
		sourceType, sourceID,
	)
	return err
}

func (r *ragRepository) DeleteAllBySourceType(ctx context.Context, sourceType string) error {
	_, err := r.DB.ExecContext(ctx,
		`DELETE FROM rag_chunks WHERE source_type = $1`,
		sourceType,
	)
	return err
}

func (r *ragRepository) GetCarsContentForIndexing(ctx context.Context) ([]CarContentRow, error) {
	query := `
		SELECT c.id AS car_id,
			COALESCE(TRIM(
				COALESCE(m.name, '') || ' ' || COALESCE(mo.name, '') || '. ' ||
				COALESCE(c.ref_no::text, '') || ' ' ||
				COALESCE(c.package, '') || ' ' ||
				COALESCE(c.body_type::text, '') || ' ' ||
				COALESCE(c.year::text, '') || ' ' ||
				COALESCE(c.color, '') || ' ' ||
				COALESCE(c.fuel::text, '') || ' ' ||
				COALESCE(c.transmission::text, '') || ' ' ||
				COALESCE(c.drive::text, '') || ' ' ||
				COALESCE(cd.full_title, '') || ' ' ||
				COALESCE(cd.description, '') || ' ' ||
				COALESCE((SELECT string_agg(COALESCE(csd.title, '') || ' ' || COALESCE(csd.description, ''), ' ')
				 FROM car_sub_details csd WHERE csd.car_detail_id = cd.id), '')
			), '') AS content
		FROM cars c
		JOIN car_models mo ON mo.id = c.model_id
		JOIN car_makes m ON m.id = mo.make_id
		LEFT JOIN car_details cd ON cd.car_id = c.id
		ORDER BY c.id
	`
	var rows []CarContentRow
	err := r.DB.SelectContext(ctx, &rows, query)
	return rows, err
}
