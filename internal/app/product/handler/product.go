package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/valentinusdelvin/savebite-be/internal/app/product/usecase"
	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
	"github.com/valentinusdelvin/savebite-be/internal/models"
	"gorm.io/gorm"
)

type ProductHandler struct {
	validator      *validator.Validate
	productUsecase usecase.ProductUsecaseItf
}

func NewProductHandler(routerGroup *gin.RouterGroup, validator *validator.Validate, productUsecase usecase.ProductUsecaseItf) {
	ProductHandler := ProductHandler{
		validator:      validator,
		productUsecase: productUsecase,
	}
	product := routerGroup.Group("/products")

	product.POST("/", ProductHandler.CreateProduct)
	product.GET("/", ProductHandler.GetAllProducts)
	product.GET("/:id", ProductHandler.GetProductByID)

}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var product dto.CreateProduct

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "bad request",
			Details: err.Error(),
		})
		return
	}

	if err := h.productUsecase.CreateProduct(product); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "internal server error",
			Details: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONSuccessResponse{
		Status:  http.StatusCreated,
		Message: "product created successfully",
		Data:    product,
	})
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := h.productUsecase.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "internal server error",
			Details: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, models.JSONSuccessResponse{
		Status:  http.StatusOK,
		Message: "products retrieved successfully",
		Data:    products,
	})
}

func (h *ProductHandler) GetProductByID(ctx *gin.Context) {
	var param dto.GetProductById

	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "bad request",
			Details: err.Error(),
		})
	}

	product, err := h.productUsecase.GetProductByID(param.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, models.JSONErrorResponse{
				Status:  http.StatusNotFound,
				Error:   "product not found",
				Details: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "internal server error",
			Details: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.JSONSuccessResponse{
		Status:  http.StatusOK,
		Message: "product retrieved successfully",
		Data:    product,
	})
}
