package product_usecase

import (
	"context"

	product_response "simple-golang-database/internal/modules/product/model/response"
	product_repository "simple-golang-database/internal/modules/product/repository"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, name string, price int64) (product_response.Product, error)
	GetProductByID(ctx context.Context, id string) (product_response.Product, error)
	GetAllProduct(ctx context.Context) ([]product_response.Product, error)
	UpdateProductByID(ctx context.Context, id string, name string, price int64) error
	DeleteProductByID(ctx context.Context, id string) error
}

type productUsecase struct {
	productRepository product_repository.ProductRepository
}
