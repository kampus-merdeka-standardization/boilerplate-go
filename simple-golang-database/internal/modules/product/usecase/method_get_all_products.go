package product_usecase

import (
	"context"
	product_entity "simple-golang-database/internal/modules/product/model/entity"
)

func (productUsecase *productUsecase) GetAllProduct(ctx context.Context) ([]product_entity.Product, error) {
	products, err := productUsecase.productRepository.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
