package service

import (
	"context"
	"errors"

	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
	"github.com/valentinusdelvin/savebite-be/internal/pkg/gemini"
	"google.golang.org/genai"
)

type AIServiceItf interface {
	GenerateRecipe(ctx context.Context, req dto.AIRequest) (string, error)
}

type aiService struct {
	client *genai.Client
	model  string
}

func NewAIService(client *genai.Client) AIServiceItf {
	return &aiService{
		client: client,
		model:  "gemini-2.5-flash",
	}
}

func (s *aiService) GenerateRecipe(ctx context.Context, req dto.AIRequest) (string, error) {
	prompt, err := gemini.FormatSaveBitePrompt(req)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Models.GenerateContent(
		ctx,
		s.model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 ||
		resp.Candidates[0].Content == nil ||
		len(resp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("empty response from gemini")
	}

	part := resp.Candidates[0].Content.Parts[0]

	if part.Text == "" {
		return "", errors.New("gemini response part has no text")
	}

	return part.Text, nil
}
