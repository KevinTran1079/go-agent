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
			t.Chdir(t.TempDir())
			file := "example.txt"
			if err := os.WriteFile(file, []byte(tt.contents), 0600); err != nil {
				t.Fatalf("create test file: %v", err)
			}

			got, err := ReadFile(file)
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
		file string
	}{
		{
			name: "missing file",
			file: "missing.txt",
		},
		{
			name: "missing parent directory",
			file: filepath.Join("missing", "example.txt"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Chdir(t.TempDir())
			got, err := ReadFile(tt.file)
			if !errors.Is(err, os.ErrNotExist) {
				t.Errorf("ReadFile() error = %v; want an error wrapping os.ErrNotExist", err)
			}
			if got != "" {
				t.Errorf("ReadFile() = %q; want empty string", got)
			}
		})
	}
}
