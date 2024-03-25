package product_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	product_model "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/model/entity"
)

type Repository interface {
	CreateProduct(ctx context.Context, tx *sqlx.Tx, name string, price float64) (id string, err error)
	GetProductByID(ctx context.Context, tx *sqlx.Tx, id string) (*product_model.Product, error)
	GetAllProduct(ctx context.Context, tx *sqlx.Tx) ([]product_model.Product, error)
	UpdateProductByID(ctx context.Context, tx *sqlx.Tx, id string, name string, price float64) error
	DeleteProductByID(ctx context.Context, tx *sqlx.Tx, id string) error
}
