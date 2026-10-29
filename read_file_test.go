package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadFileExistingFile(t *testing.T) {
	tests := []struct {
		name     string
		contents string
	}{
		{
			name:     "text file",
			contents: "hello, agent\n",
		},
		{
			name:     "text file no content",
			contents: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "example.txt")
			if err := os.WriteFile(path, []byte(tt.contents), 0600); err != nil {
				t.Fatalf("create test file: %v", err)
			}

			got, err := ReadFile(path)
			if err != nil {
				t.Fatalf("ReadFile() returned unexpected error: %v", err)
			}
			if got != tt.contents {
				t.Errorf("ReadFile() = %q; want %q", got, tt.contents)
			}
		})
	}
}

func TestReadFileNotFound(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{
			name: "missing file",
			path: "missing.txt",
		},
		{
			name: "missing parent directory",
			path: filepath.Join("missing", "example.txt"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), tt.path)
			got, err := ReadFile(path)
			if !errors.Is(err, os.ErrNotExist) {
				t.Errorf("ReadFile() error = %v; want an error wrapping os.ErrNotExist", err)
			}
			if got != "" {
				t.Errorf("ReadFile() = %q; want empty string", got)
			}
		})
	}
}
