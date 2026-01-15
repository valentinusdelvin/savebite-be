package repository

import (
	"github.com/valentinusdelvin/savebite-be/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductRepositoryItf interface {
	GetProductTest() string
	CreateProduct(param entity.Product) error
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(id string) (entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryItf {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetProductTest() string {
	return "Product Repository Test"
}

func (r *productRepository) CreateProduct(param entity.Product) error {
	err := r.db.Table("products").Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Table("products").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProductByID(id string) (entity.Product, error) {
	var product entity.Product
	err := r.db.Table("products").Where("id = ?", id).First(&product).Error
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}
