package rag

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/user/car-project/internal/repository"
)

const (
	SourceTypeCar = "car"
)

// AskResult is the response from RAG Ask.
type AskResult struct {
	Answer   string   `json:"answer"`
	Sources  []Source `json:"sources,omitempty"`
}

// Source references a chunk that was used.
type Source struct {
	SourceType string `json:"source_type"`
	SourceID   string `json:"source_id"`
	Content    string `json:"content"`
}

// RAG orchestrates retrieval and generation.
type RAG struct {
	embedder Embedder
	llm      LLM
	repo     repository.RAGRepository
	topK     int
}

// NewRAG creates a RAG pipeline.
func NewRAG(embedder Embedder, llm LLM, repo repository.RAGRepository, topK int) *RAG {
	if topK <= 0 {
		topK = 5
	}
	return &RAG{embedder: embedder, llm: llm, repo: repo, topK: topK}
}

// Ask runs retrieval-augmented generation: embed query -> search -> generate answer.
func (r *RAG) Ask(ctx context.Context, query string) (*AskResult, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return &AskResult{Answer: "Please provide a question."}, nil
	}

	emb, err := r.embedder.Embed(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("embed query: %w", err)
	}
	embeddingStr := FormatVectorForPG(emb)

	chunks, err := r.repo.SearchByEmbedding(ctx, embeddingStr, r.topK)
	if err != nil {
		return nil, fmt.Errorf("search: %w", err)
	}

	sources := make([]Source, 0, len(chunks))
	var contextParts []string
	for _, c := range chunks {
		contextParts = append(contextParts, c.Content)
		sources = append(sources, Source{
			SourceType: c.SourceType,
			SourceID:   c.SourceID,
			Content:    truncate(c.Content, 200),
		})
	}

	contextBlob := strings.Join(contextParts, "\n\n---\n\n")
	if contextBlob == "" {
		contextBlob = "No relevant documents found in the knowledge base."
	}

	systemPrompt := `You are a helpful assistant for a car dealership/inventory system. Answer the user's question using ONLY the provided context. If the context does not contain enough information, say so. Be concise and factual. Do not make up car details or inventory.`
	userMessage := fmt.Sprintf("Context:\n%s\n\nQuestion: %s", contextBlob, query)

	answer, err := r.llm.Complete(ctx, systemPrompt, userMessage)
	if err != nil {
		return nil, fmt.Errorf("generate: %w", err)
	}

	return &AskResult{
		Answer:  answer,
		Sources: sources,
	}, nil
}

// IndexCars loads all car content from the DB, embeds it, and upserts into rag_chunks.
func (r *RAG) IndexCars(ctx context.Context) (indexed int, err error) {
	rows, err := r.repo.GetCarsContentForIndexing(ctx)
	if err != nil {
		return 0, fmt.Errorf("get cars content: %w", err)
	}

	if err := r.repo.DeleteAllBySourceType(ctx, SourceTypeCar); err != nil {
		return 0, fmt.Errorf("clear existing car chunks: %w", err)
	}

	const maxChunkSize = 1500
	var chunks []repository.RAGChunkInput
	for _, row := range rows {
		if strings.TrimSpace(row.Content) == "" {
			continue
		}
		content := normalizeSpace(row.Content)
		sourceID := strconv.FormatInt(row.CarID, 10)
		parts := splitIntoChunks(content, maxChunkSize)
		for _, part := range parts {
			emb, err := r.embedder.Embed(ctx, part)
			if err != nil {
				return 0, fmt.Errorf("embed car %s: %w", sourceID, err)
			}
			meta, _ := json.Marshal(map[string]interface{}{"car_id": row.CarID})
			chunks = append(chunks, repository.RAGChunkInput{
				SourceType: SourceTypeCar,
				SourceID:   sourceID,
				Content:    part,
				Embedding:  FormatVectorForPG(emb),
				Metadata:   string(meta),
			})
		}
		indexed++
	}

	const batchSize = 20
	for i := 0; i < len(chunks); i += batchSize {
		end := i + batchSize
		if end > len(chunks) {
			end = len(chunks)
		}
		batch := chunks[i:end]
		if err := r.repo.UpsertChunks(ctx, batch); err != nil {
			return 0, fmt.Errorf("upsert chunks: %w", err)
		}
	}
	return indexed, nil
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func normalizeSpace(s string) string {
	return strings.TrimSpace(strings.Join(strings.Fields(s), " "))
}

func splitIntoChunks(s string, maxSize int) []string {
	if len(s) <= maxSize {
		return []string{s}
	}
	var out []string
	for len(s) > 0 {
		end := maxSize
		if end > len(s) {
			end = len(s)
		} else {
			// Try to break at space
			if idx := strings.LastIndex(s[:end], " "); idx > maxSize/2 {
				end = idx + 1
			}
		}
		out = append(out, strings.TrimSpace(s[:end]))
		s = strings.TrimSpace(s[end:])
	}
	return out
}
