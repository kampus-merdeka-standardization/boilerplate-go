package product_postgres

import (
	"github.com/jmoiron/sqlx"
	product_repository "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/repository"
)

func NewProductRepository(db *sqlx.DB) product_repository.ProductRepository {
	return &productPostgresRepository{
		db: db,
	}
}
