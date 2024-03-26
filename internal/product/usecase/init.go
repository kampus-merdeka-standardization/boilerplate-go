package product_usecase

import (
	"github.com/jmoiron/sqlx"
	product_repository "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/repository"
)

func NewProductUsecase(db *sqlx.DB, productRepository product_repository.ProductRepository) ProductUsecase {
	return &productUsecase{}
}
