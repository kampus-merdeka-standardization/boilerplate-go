package product_usecase

import (
	"context"

	product_entity "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/model/entity"
)

func (productUsecase *productUsecase) GetProductByID(ctx context.Context, id string) (*product_entity.Product, error) {
	product, err := productUsecase.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
