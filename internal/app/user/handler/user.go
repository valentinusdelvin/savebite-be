package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/valentinusdelvin/savebite-be/internal/app/user/usecase"
	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
	"github.com/valentinusdelvin/savebite-be/internal/models"
	"gorm.io/gorm"
)

type userHandler struct {
	validator   *validator.Validate
	UserUsecase usecase.UserUsecaseItf
}

func NewUserHandler(routerGroup *gin.RouterGroup, validator *validator.Validate, UserUsecase usecase.UserUsecaseItf) {
	UserHandler := userHandler{
		validator:   validator,
		UserUsecase: UserUsecase,
	}
	user := routerGroup.Group("/users")

	user.POST("/register", UserHandler.Register)
	user.POST("/login", UserHandler.Login)
}

func (h *userHandler) Register(ctx *gin.Context) {
	param := dto.Register{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "Bad Request",
			Details: err.Error(),
		})
		return
	}

	err = h.validator.Struct(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "Bad Request",
			Details: err.Error(),
		})
		return
	}

	err = h.UserUsecase.Register(param)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(http.StatusConflict, models.JSONErrorResponse{
				Status:  http.StatusConflict,
				Error:   "Conflict",
				Details: "Email already exists",
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
				Status:  http.StatusInternalServerError,
				Error:   "Internal Server Error",
				Details: err.Error(),
			})
			return
		}
	}
}

func (h *userHandler) Login(ctx *gin.Context) {
	param := dto.Login{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "Bad Request",
			Details: err.Error(),
		})
		return
	}

	token, err := h.UserUsecase.Login(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Internal Server Error",
			Details: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONSuccessResponse{
		Status:  http.StatusOK,
		Message: "login successful",
		Data:    token,
	})
}
