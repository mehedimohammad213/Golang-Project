package service

import (
	"context"

	"github.com/user/car-project/internal/rag"
)

type RAGService interface {
	Ask(ctx context.Context, query string) (*rag.AskResult, error)
	IndexCars(ctx context.Context) (int, error)
}

type ragService struct {
	rag *rag.RAG
}

func NewRAGService(r *rag.RAG) RAGService {
	return &ragService{rag: r}
}

func (s *ragService) Ask(ctx context.Context, query string) (*rag.AskResult, error) {
	return s.rag.Ask(ctx, query)
}

func (s *ragService) IndexCars(ctx context.Context) (int, error) {
	return s.rag.IndexCars(ctx)
}
