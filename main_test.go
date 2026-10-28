package main

import (
	"context"
	"testing"

	"google.golang.org/genai"
)

type fakeModel struct {
	response *genai.GenerateContentResponse
}

func (m *fakeModel) GenerateContent(context.Context, []*genai.Content) (*genai.GenerateContentResponse, error) {
	return m.response, nil
}

func TestRunInvalidResponse(t *testing.T) {
	tests := []struct {
		name     string
		response *genai.GenerateContentResponse
		wantErr  string
	}{
		{
			name:     "nil candidates",
			response: &genai.GenerateContentResponse{},
			wantErr:  "response has no candidates",
		},
		{
			name: "empty candidates",
			response: &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{},
			},
			wantErr: "response has no candidates",
		},
		{
			name: "candidate without content",
			response: &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{{}},
			},
			wantErr: "response candidate content has no content",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := &fakeModel{response: tt.response}
			err := Run(context.Background(), nil, model)
			if err == nil {
				t.Fatalf("Run() returned nil; want error %q", tt.wantErr)
			}
			if err.Error() != tt.wantErr {
				t.Errorf("Run() error = %q; want %q", err.Error(), tt.wantErr)
			}
		})
	}
}
