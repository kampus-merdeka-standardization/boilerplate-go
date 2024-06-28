package product_usecase

import (
	"context"
	"database/sql"
	product_response "simple-golang-database/internal/modules/product/model/response"
	pkg_error "simple-golang-database/pkg/error"
)

func (productUsecase *productUsecase) GetProductByID(ctx context.Context, id string) (product_response.Product, error) {
	product, err := productUsecase.productRepository.GetProductByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return product_response.Product{}, pkg_error.NewNotFound(err, "Product Not found")
		}
		return product_response.Product{}, err
	}

	return product_response.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
