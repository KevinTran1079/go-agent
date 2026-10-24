package main

import (
	"context"
	"fmt"

	"os"

	"github.com/KevinTran1079/go-agent/internal/gemini"
	"github.com/joho/godotenv"
)

func main() {
	Run("describe what an llm does in 5 sentences")
}

func Run(prompt string) error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("could not load .env: %w", err)
	}
	ctx := context.Background()
	model, err := gemini.NewClient(
		ctx,
		os.Getenv("GEMINI_API_KEY"),
		"gemini-3.6-flash",
	)
	if err != nil {
		return fmt.Errorf("could not create model client: %w", err)
	}

	resp, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		return fmt.Errorf("could not generate content: %w", err)
	}

	fmt.Println(resp)
	return nil
}
