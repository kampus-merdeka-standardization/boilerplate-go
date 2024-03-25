package product_postgres

import product_repository "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/repository"

func NewProductRepository() product_repository.ProductRepository {
	return &productPostgresRepository{}
}
