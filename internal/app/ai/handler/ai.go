package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/savebite-be/internal/app/ai/usecase"
	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
)

type AIHandler struct {
	aiUsecase usecase.AiUsecaseItf
}

func NewAIHandler(routerGroup *gin.RouterGroup, aiUsecase usecase.AiUsecaseItf) {
	AIHandler := AIHandler{
		aiUsecase: aiUsecase,
	}
	ai := routerGroup.Group("/ai")

	ai.POST("/recipes", AIHandler.GenerateRecipe)
}

func (h *AIHandler) GenerateRecipe(c *gin.Context) {
	var req dto.AIRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	result, err := h.aiUsecase.GenerateRecipe(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to generate recipe",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"recipe": result,
		},
	})
}
