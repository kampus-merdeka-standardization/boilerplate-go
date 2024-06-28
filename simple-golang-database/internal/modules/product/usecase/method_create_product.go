package product_usecase

import (
	"context"
	product_response "simple-golang-database/internal/modules/product/model/response"
)

func (productUsecase *productUsecase) CreateProduct(ctx context.Context, name string, price int64) (product_response.Product, error) {
	product, err := productUsecase.productRepository.CreateProduct(ctx, name, price)
	if err != nil {
		return product_response.Product{}, err
	}

	return product_response.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
