package usecase

import (
	"github.com/google/uuid"
	"github.com/valentinusdelvin/savebite-be/internal/app/product/repository"
	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
	"github.com/valentinusdelvin/savebite-be/internal/domain/entity"
)

type ProductUsecaseItf interface {
	GetProductTest() string
	CreateProduct(param dto.CreateProduct) error
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(id string) (entity.Product, error)
}
type ProductUsecase struct {
	productRepo repository.ProductRepositoryItf
}

func NewProductUsecase(productRepo repository.ProductRepositoryItf) ProductUsecaseItf {
	return &ProductUsecase{
		productRepo: productRepo,
	}
}

func (u *ProductUsecase) GetProductTest() string {
	return u.productRepo.GetProductTest()
}

func (u *ProductUsecase) CreateProduct(param dto.CreateProduct) error {
	product := entity.Product{
		ProductId:    uuid.New(),
		Name:         param.Name,
		Price:        param.Price,
		Description:  param.Description,
		Stock:        param.Stock,
		ProductImage: param.ProductImage,
	}

	err := u.productRepo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductUsecase) GetAllProducts() ([]entity.Product, error) {
	products, err := u.productRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (u *ProductUsecase) GetProductByID(id string) (entity.Product, error) {
	product, err := u.productRepo.GetProductByID(id)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}
