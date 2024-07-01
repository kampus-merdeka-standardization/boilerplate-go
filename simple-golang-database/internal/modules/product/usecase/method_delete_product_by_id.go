package product_usecase

import (
	"context"
)

func (productUsecase *productUsecase) DeleteProductByID(ctx context.Context, id string) error {
	err := productUsecase.productRepository.DeleteProductByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
