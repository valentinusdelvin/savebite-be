package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/savebite-be/internal/app/ai/usecase"
	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
	"github.com/valentinusdelvin/savebite-be/internal/models"
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
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "bad request",
			Details: err.Error(),
		})
		return
	}

	result, err := h.aiUsecase.GenerateRecipe(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "failed to generate recipe",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONSuccessResponse{
		Status:  http.StatusOK,
		Message: "recipe generated successfully",
		Data:    result,
	})
}
