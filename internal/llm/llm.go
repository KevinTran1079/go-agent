package llm

import (
	"context"

	"google.golang.org/genai"
)

type LLM interface {
	GenerateContent(ctx context.Context, prompt []*genai.Content) (*genai.GenerateContentResponse, error)
}
