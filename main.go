package main

import (
	"bufio"
	"context"
	"fmt"

	"os"

	"github.com/KevinTran1079/go-agent/internal/gemini"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

type Message struct {
	Role    string
	Content string
}

func main() {
	var messages []*genai.Content
	messages = append(messages, genai.NewContentFromText("This is my first message through the Gemini Go SDK", genai.RoleUser))
	err := Run(messages)
	if err != nil {
		fmt.Println("error running main %w", err)
	}

}

func Run(messages []*genai.Content) error {
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
