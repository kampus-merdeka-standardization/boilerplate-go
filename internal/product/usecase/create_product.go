package product_usecase

import (
	"context"
)

func (productUsecase *productUsecase) CreateProduct(ctx context.Context, name string, price float64) (string, error) {
	id, err := productUsecase.productRepository.CreateProduct(ctx, name, price)
	if err != nil {
		return "", err
	}

	return id, nil
}
