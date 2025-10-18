package services

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiService struct {
	client *genai.Client
	ctx    context.Context
}

func NewGeminiService(ctx context.Context) (*GeminiService, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY not configured")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &GeminiService{
		client: client,
		ctx:    ctx,
	}, nil
}

func (g *GeminiService) Generate(prompt string) (string, error) {
	model := g.client.GenerativeModel("gemini-2.5-flash")

	resp, err := model.GenerateContent(g.ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response generated")
	}

	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}

func (g *GeminiService) Close() {
	if g.client != nil {
		g.client.Close()
	}
}
