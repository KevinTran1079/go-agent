package main

import (
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
