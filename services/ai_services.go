package services

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func InitAIService(ctx context.Context, apiKey string) (*genai.GenerativeModel, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %v", err)
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	return model, nil
}
