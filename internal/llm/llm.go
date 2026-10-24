package llm

import "context"

type LLM interface {
	Generator(ctx context.Context, prompt string) (string, error)
}
