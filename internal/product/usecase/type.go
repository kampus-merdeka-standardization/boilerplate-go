package product_usecase

import (
	"context"

	"github.com/jmoiron/sqlx"
	product_entity "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/model/entity"
	product_repository "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/repository"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, name string, price float64) (string, error)
	GetProductByID(ctx context.Context, id string) (*product_entity.Product, error)
	GetAllProduct(ctx context.Context) ([]product_entity.Product, error)
	UpdateProductByID(ctx context.Context, id string, name string, price float64) error
	DeleteProductByID(ctx context.Context, id string) error
}

type productUsecase struct {
	db                *sqlx.DB
	productRepository product_repository.ProductRepository
}
