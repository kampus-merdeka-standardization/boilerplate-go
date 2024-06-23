package product_usecase

import (
	"context"

	product_entity "simple-golang-database/internal/modules/product/model/entity"
	product_repository "simple-golang-database/internal/modules/product/repository"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, name string, price float64) (string, error)
	GetProductByID(ctx context.Context, id string) (*product_entity.Product, error)
	GetAllProduct(ctx context.Context) ([]product_entity.Product, error)
	UpdateProductByID(ctx context.Context, id string, name string, price float64) error
	DeleteProductByID(ctx context.Context, id string) error
}

type productUsecase struct {
	productRepository product_repository.ProductRepository
}
