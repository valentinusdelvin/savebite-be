package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/savebite-be/internal/app/product/usecase"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecaseItf
}

func NewProductHandler(routerGroup *gin.RouterGroup, productUsecase usecase.ProductUsecaseItf) {
	ProductHandler := ProductHandler{
		productUsecase: productUsecase,
	}
	product := routerGroup.Group("/products")

	product.GET("/", ProductHandler.GetAllProducts)

}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	product := h.productUsecase.GetProductTest()

	ctx.JSON(200, gin.H{
		"message": product,
	})
}
