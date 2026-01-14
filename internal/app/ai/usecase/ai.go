package usecase

import (
	"context"
	"encoding/json"

	"github.com/valentinusdelvin/savebite-be/internal/app/ai/service"
	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
)

type AiUsecaseItf interface {
	GenerateRecipe(ctx context.Context, req dto.AIRequest) (dto.Recipe, error)
}

type AiUsecase struct {
	aiService service.AIService
}

func NewAiUsecase(ai service.AIService) *AiUsecase {
	return &AiUsecase{aiService: ai}
}

func (a *AiUsecase) GenerateRecipe(ctx context.Context, req dto.AIRequest) (dto.Recipe, error) {
	result, err := a.aiService.GenerateRecipe(ctx, req)
	if err != nil {
		return dto.Recipe{}, err
	}

	var recipe dto.Recipe
	if err := json.Unmarshal([]byte(result), &recipe); err != nil {
		return dto.Recipe{}, err
	}

	return recipe, nil
}
