package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/savebite-be/internal/app/user/usecase"
	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecaseItf
}

func NewUserHandler(routerGroup *gin.RouterGroup, UserUsecase usecase.UserUsecaseItf) {
	UserHandler := UserHandler{
		UserUsecase: UserUsecase,
	}
	user := routerGroup.Group("/users")

	user.POST("/register", UserHandler.Register)
	user.POST("/login", UserHandler.Login)
}

func (h *UserHandler) Register(ctx *gin.Context) {
	param := dto.Register{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = h.UserUsecase.Register(param)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(409, gin.H{
				"error":   "Conflict",
				"message": "Email already exists",
			})
			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
			return
		}
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	param := dto.Login{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	token, err := h.UserUsecase.Login(param)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}
