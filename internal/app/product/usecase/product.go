package usecase

import "github.com/valentinusdelvin/savebite-be/internal/app/product/repository"

type ProductUsecaseItf interface {
	GetProductTest() string
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
