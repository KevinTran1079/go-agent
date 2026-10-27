package main

import (
	"bufio"
	"context"
	"fmt"

	"os"

	"github.com/KevinTran1079/go-agent/internal/gemini"
	"github.com/KevinTran1079/go-agent/internal/llm"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

type Message struct {
	Role    string
	Content string
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("unable to load .env")
		os.Exit(1)
	}

	ctx := context.Background()
	model, err := gemini.NewClient(
		ctx,
		os.Getenv("GEMINI_API_KEY"),
		"gemini-3.6-flash",
	)
	if err != nil {
		fmt.Println("could not create model client: %w", err)
	}

	var messages []*genai.Content
	messages = append(messages, genai.NewContentFromText("This is my first message through the Gemini Go SDK", genai.RoleUser))

	err = Run(ctx, messages, model)
	if err != nil {
		fmt.Println("error running main %w", err)
	}

}

func Run(ctx context.Context, messages []*genai.Content, model llm.LLM) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		resp, err := model.GenerateContent(ctx, messages)
		if err != nil {
			return fmt.Errorf("could not generate content: %w", err)
		}
		messages = append(messages, resp.Candidates[0].Content)
		fmt.Println("model:", resp.Text())
		fmt.Print("you: ")

		if !scanner.Scan() {
			return scanner.Err()
		}

		input := scanner.Text()
		if input == "exit" {
			return nil
		}
		messages = append(messages, genai.NewContentFromText(input, genai.RoleUser))
	}
}
