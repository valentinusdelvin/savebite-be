package dto

type CreateProduct struct {
	Name         string  `form:"name" binding:"required, min=3"`
	Description  string  `form:"description" binding:"required, min=10"`
	Price        float64 `form:"price" binding:"required"`
	Stock        int     `form:"stock"`
	ProductImage string  `form:"product_image"`
}

type GetProductById struct {
	ProductId string `json:"product_id" binding:"required"`
}
