package product_usecase

import (
	"context"

	product_entity "simple-golang-database/internal/modules/product/model/entity"
)

func (productUsecase *productUsecase) CreateProduct(ctx context.Context, name string, price float64) (string, error) {
	id, err := productUsecase.productRepository.CreateProduct(ctx, name, price)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (productUsecase *productUsecase) GetProductByID(ctx context.Context, id string) (*product_entity.Product, error) {
	product, err := productUsecase.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (productUsecase *productUsecase) GetAllProduct(ctx context.Context) ([]product_entity.Product, error) {
	products, err := productUsecase.productRepository.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

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
