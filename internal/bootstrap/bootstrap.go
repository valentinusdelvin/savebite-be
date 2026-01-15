package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	aiHandler "github.com/valentinusdelvin/savebite-be/internal/app/ai/handler"
	aiService "github.com/valentinusdelvin/savebite-be/internal/app/ai/service"
	aiUsecase "github.com/valentinusdelvin/savebite-be/internal/app/ai/usecase"
	productHandler "github.com/valentinusdelvin/savebite-be/internal/app/product/handler"
	productrepository "github.com/valentinusdelvin/savebite-be/internal/app/product/repository"
	productusecase "github.com/valentinusdelvin/savebite-be/internal/app/product/usecase"
	userHandler "github.com/valentinusdelvin/savebite-be/internal/app/user/handler"
	userRepository "github.com/valentinusdelvin/savebite-be/internal/app/user/repository"
	userUsecase "github.com/valentinusdelvin/savebite-be/internal/app/user/usecase"
	"github.com/valentinusdelvin/savebite-be/internal/infra/config"
	"github.com/valentinusdelvin/savebite-be/internal/infra/mysql"
	"github.com/valentinusdelvin/savebite-be/internal/middleware"
	"github.com/valentinusdelvin/savebite-be/internal/pkg/jwt"
	"google.golang.org/genai"
)

func Start() error {
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		return err
	}

	database, err := mysql.New(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_DATABASE))

	if err != nil {
		return err
	}

	if err := mysql.Migrate(database); err != nil {
		return err
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.GEMINI_API_KEY,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return err
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")

	jwtItf := jwt.NewJWT(cfg.JWT_SECRET, cfg.JWT_EXPIRES)

	userRepo := userRepository.NewUserRepository(database)
	userUsecase := userUsecase.NewUserUsecase(userRepo, jwtItf)
	userHandler.NewUserHandler(v1, userUsecase)

	v1.Use(middleware.NewMiddleware(jwtItf).Authentication)

	productRepo := productrepository.NewProductRepository(database)
	productUsecase := productusecase.NewProductUsecase(productRepo)
	productHandler.NewProductHandler(v1, productUsecase)

	aiService := aiService.NewAIService(client)
	aiUsecase := aiUsecase.NewAiUsecase(aiService)
	aiHandler.NewAIHandler(v1, aiUsecase)

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err = r.Run()
	if err != nil {
		return err
	}

	return nil
}
