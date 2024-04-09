package product_usecase

import "context"

func (productUsecase *productUsecase) DeleteProductByID(ctx context.Context, id string) error {
	product, err := productUsecase.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return err
	}

	err = productUsecase.productRepository.DeleteProductByID(ctx, product.ID)
	if err != nil {
		return err
	}

	return nil
}
