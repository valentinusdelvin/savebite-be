package repository

import "gorm.io/gorm"

type ProductRepositoryItf interface {
	GetProductTest() string
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryItf {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProductTest() string {
	return "Product Repository Test"
}
