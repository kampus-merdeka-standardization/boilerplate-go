package product_usecase

import (
	product_repository "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/repository"
)

func NewProductUsecase(productRepository product_repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepository: productRepository}
}
