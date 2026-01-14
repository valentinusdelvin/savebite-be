package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/savebite-be/internal/app/product/usecase"
	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
	"gorm.io/gorm"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecaseItf
}

func NewProductHandler(routerGroup *gin.RouterGroup, productUsecase usecase.ProductUsecaseItf) {
	ProductHandler := ProductHandler{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	if err := h.productUsecase.CreateProduct(product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": product,
	})
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := h.productUsecase.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": products,
	})
}

func (h *ProductHandler) GetProductByID(ctx *gin.Context) {
	var param dto.GetProductById

	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
	}

	product, err := h.productUsecase.GetProductByID(param.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "product not found",
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
