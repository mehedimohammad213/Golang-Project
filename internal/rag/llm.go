package rag

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

// LLM generates text completions (for RAG answer generation).
type LLM interface {
	Complete(ctx context.Context, systemPrompt, userMessage string) (string, error)
}

type openAILLM struct {
	client *openai.Client
	model  string
}

// NewOpenAILLM creates an LLM using OpenAI's Chat API.
func NewOpenAILLM(apiKey, model string) LLM {
	if model == "" {
		model = openai.GPT4oMini
	}
	return &openAILLM{
		client: openai.NewClient(apiKey),
		model:  model,
	}
}

func (l *openAILLM) Complete(ctx context.Context, systemPrompt, userMessage string) (string, error) {
	resp, err := l.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: l.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: userMessage},
		},
		MaxTokens: 1024,
	})
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no completion returned")
	}
	return resp.Choices[0].Message.Content, nil
}
