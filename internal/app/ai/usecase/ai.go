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

type aiUsecase struct {
	aiService service.AIServiceItf
}

func NewAiUsecase(aiService service.AIServiceItf) AiUsecaseItf {
	return &aiUsecase{aiService: aiService}
}

func (a *aiUsecase) GenerateRecipe(ctx context.Context, req dto.AIRequest) (dto.Recipe, error) {
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
