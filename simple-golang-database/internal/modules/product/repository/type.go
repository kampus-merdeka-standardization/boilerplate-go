package product_repository

import (
	"context"

	product_entity "simple-golang-database/internal/modules/product/model/entity"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, name string, price int64) (product product_entity.Product, err error)
	GetProductByID(ctx context.Context, id string) (product_entity.Product, error)
	GetAllProduct(ctx context.Context) ([]product_entity.Product, error)
	UpdateProductByID(ctx context.Context, id string, name string, price int64) (product_entity.Product, error)
	DeleteProductByID(ctx context.Context, id string) error
}
