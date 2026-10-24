package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load .env: ", err)
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:      os.Getenv("GEMINI_API_KEY"),
		Backend:     genai.BackendGeminiAPI,
		HTTPOptions: genai.HTTPOptions{APIVersion: "v1"},
	})
	if err != nil {
		fmt.Println("Could not create client.")
		os.Exit(1)
	}
	resp, err := client.Models.GenerateContent(ctx,
		"gemini-3.6-flash",
		genai.Text("How does AI work in 5 sentences"),
		nil,
	)
	if err != nil {
		fmt.Println("Failed to generate text: %w", err)
		os.Exit(1)
	}
	fmt.Println(resp.Text())

}
