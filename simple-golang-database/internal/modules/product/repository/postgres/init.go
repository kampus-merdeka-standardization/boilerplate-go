package product_postgres

import (
	product_repository "simple-golang-database/internal/modules/product/repository"

	"github.com/jmoiron/sqlx"
)

func NewProductRepository(db *sqlx.DB) product_repository.ProductRepository {
	return &productPostgresRepository{
		db: db,
	}
}
