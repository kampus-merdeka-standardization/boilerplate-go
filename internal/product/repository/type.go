package product_repository

import (
	"context"

	product_model "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/model/entity"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, name string, price float64) (id string, err error)
	GetProductByID(ctx context.Context, id string) (*product_model.Product, error)
	GetAllProduct(ctx context.Context) ([]product_model.Product, error)
	UpdateProductByID(ctx context.Context, id string, name string, price float64) error
	DeleteProductByID(ctx context.Context, id string) error
}
