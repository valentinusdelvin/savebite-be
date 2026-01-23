package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	aiHandler "github.com/valentinusdelvin/savebite-be/internal/app/ai/handler"
	aiService "github.com/valentinusdelvin/savebite-be/internal/app/ai/service"
	aiUsecase "github.com/valentinusdelvin/savebite-be/internal/app/ai/usecase"
	productHandler "github.com/valentinusdelvin/savebite-be/internal/app/product/handler"
	productRepository "github.com/valentinusdelvin/savebite-be/internal/app/product/repository"
	productUsecase "github.com/valentinusdelvin/savebite-be/internal/app/product/usecase"
	userHandler "github.com/valentinusdelvin/savebite-be/internal/app/user/handler"
	userRepository "github.com/valentinusdelvin/savebite-be/internal/app/user/repository"
	userUsecase "github.com/valentinusdelvin/savebite-be/internal/app/user/usecase"
	"github.com/valentinusdelvin/savebite-be/internal/infra/config"
	postgres "github.com/valentinusdelvin/savebite-be/internal/infra/postgresql"
	"github.com/valentinusdelvin/savebite-be/internal/middleware"
	"github.com/valentinusdelvin/savebite-be/internal/models"
	"github.com/valentinusdelvin/savebite-be/internal/pkg/jwt"
	"google.golang.org/genai"
)

func Start() error {
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		return err
	}

	database, err := postgres.New(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DB_HOST,
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_DATABASE,
		cfg.DB_PORT,
	))

	if err != nil {
		return err
	}

	if err := postgres.Migrate(database); err != nil {
		return err
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.GEMINI_API_KEY,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return err
	}

	validator := validator.New()

	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.JSONSuccessResponse{
			Status:  http.StatusOK,
			Message: "pong",
			Data:    "pong",
		})
	})

	jwtItf := jwt.NewJWT(cfg.JWT_SECRET, cfg.JWT_EXPIRES)

	userRepo := userRepository.NewUserRepository(database)
	userUsecase := userUsecase.NewUserUsecase(userRepo, jwtItf)
	userHandler.NewUserHandler(v1, validator, userUsecase)

	v1.Use(middleware.NewMiddleware(jwtItf).Authentication)

	productRepo := productRepository.NewProductRepository(database)
	productUsecase := productUsecase.NewProductUsecase(productRepo)
	productHandler.NewProductHandler(v1, validator, productUsecase)

	aiService := aiService.NewAIService(client)
	aiUsecase := aiUsecase.NewAiUsecase(aiService)
	aiHandler.NewAIHandler(v1, aiUsecase)

	err = r.Run()
	if err != nil {
		return err
	}

	return nil
}
