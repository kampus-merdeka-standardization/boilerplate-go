package product_usecase

import "context"

func (productUsecase *productUsecase) UpdateProductByID(ctx context.Context, id string, name string, price float64) error {
	product, err := productUsecase.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return err
	}

	err = productUsecase.productRepository.UpdateProductByID(ctx, product.ID, name, price)
	if err != nil {
		return err
	}

	return nil
}
