package main

import (
	"fmt"
	"os"
)

func ReadFile(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("could not read file: %w", err)
	}
	return string(data), nil
}
