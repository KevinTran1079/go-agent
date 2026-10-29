package gemini

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type Client struct {
	client *genai.Client
	model  string
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
		client: client,
		model:  model,
	}, nil
}

func (c *Client) GenerateContent(ctx context.Context, prompt []*genai.Content) (*genai.GenerateContentResponse, error) {
	config := &genai.GenerateContentConfig{
		Tools: []*genai.Tool{
			{
				FunctionDeclarations: []*genai.FunctionDeclaration{
					{
						Name:        "read_file",
						Description: "Read the text contents of a file.",
						Parameters: &genai.Schema{
							Type: genai.TypeObject,
							Properties: map[string]*genai.Schema{
								"path": {
									Type:        genai.TypeString,
									Description: "file name to read",
								},
							},
							Required: []string{"path"},
						},
					},
				},
			},
		},
	}
	resp, err := c.client.Models.GenerateContent(
		ctx,
		c.model,
		prompt,
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("generate content: %w", err)
	}
	return resp, nil
}
