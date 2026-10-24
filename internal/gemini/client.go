package gemini

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type Client struct {
	Client *genai.Client
	Model  string
}

func NewClient(ctx context.Context, apiKey, model string) (*Client, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("create Gemini client: %w", err)
	}
	return &Client{
		Client: client,
		Model:  model,
	}, nil
}

func (c *Client) Generator(ctx context.Context, prompt string) (string, error) {
	resp, err := c.Client.Models.GenerateContent(
		ctx,
		c.Model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("generate content: %w", err)
	}
	return resp.Text(), nil
}
