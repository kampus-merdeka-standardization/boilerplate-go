package product_usecase

import (
	product_repository "simple-golang-database/internal/modules/product/repository"
)

func NewProductUsecase(productRepository product_repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepository: productRepository}
}
