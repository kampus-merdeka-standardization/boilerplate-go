package product_usecase

import (
	"context"
	"database/sql"
	product_response "simple-golang-database/internal/modules/product/model/response"
	pkg_error "simple-golang-database/pkg/error"
)

func (productUsecase *productUsecase) UpdateProductByID(ctx context.Context, id string, name string, price int64) (product_response.Product, error) {
	product, err := productUsecase.productRepository.UpdateProductByID(ctx, id, name, price)
	if err != nil {
		if err == sql.ErrNoRows {
			return product_response.Product{}, pkg_error.NewNotFound(err, "Product Not Found")
		}
		return product_response.Product{}, err
	}

	return product_response.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
