package rag

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

// Embedder generates vector embeddings for text.
type Embedder interface {
	Embed(ctx context.Context, text string) ([]float32, error)
	EmbedBatch(ctx context.Context, texts []string) ([][]float32, error)
}

type openAIEmbedder struct {
	client *openai.Client
	model  string
}

// NewOpenAIEmbedder creates an embedder using OpenAI's API.
func NewOpenAIEmbedder(apiKey, model string) Embedder {
	if model == "" {
		model = "text-embedding-3-small"
	}
	return &openAIEmbedder{
		client: openai.NewClient(apiKey),
		model:  model,
	}
}

func (e *openAIEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
	vecs, err := e.EmbedBatch(ctx, []string{text})
	if err != nil || len(vecs) == 0 {
		return nil, err
	}
	return vecs[0], nil
}

func (e *openAIEmbedder) EmbedBatch(ctx context.Context, texts []string) ([][]float32, error) {
	req := openai.EmbeddingRequestStrings{
		Input: texts,
		Model: openai.EmbeddingModel(e.model),
	}
	resp, err := e.client.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make([][]float32, len(resp.Data))
	for i, d := range resp.Data {
		out[i] = d.Embedding
	}
	return out, nil
}
