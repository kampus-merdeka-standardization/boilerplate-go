package product_usecase

import (
	"github.com/jmoiron/sqlx"
	product_repository "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/repository"
)

type ProductUsecase interface {
	CreateProduct(name string)
}

type productUsecase struct {
	db                *sqlx.DB
	productRepository product_repository.ProductRepository
}
