package usecase

import "go-api/model"

type ProductUsecase struct {
	//Repository
	repository repository.ProductRepository
}

func NewProductUseCase( repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}