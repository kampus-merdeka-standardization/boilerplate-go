package product_usecase

import (
	"context"

	product_entity "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/model/entity"
)

func (productUsecase *productUsecase) CreateProduct(ctx context.Context, name string, price float64) (string, error) {
	tx, err := productUsecase.db.Beginx()
	if err != nil {
		return "", err
	}

	id, err := productUsecase.productRepository.CreateProduct(ctx, tx, name, price)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return id, nil
}

func (productUsecase *productUsecase) GetProductByID(ctx context.Context, id string) (*product_entity.Product, error) {
	tx, err := productUsecase.db.Beginx()
	if err != nil {
		return nil, err
	}

	product, err := productUsecase.productRepository.GetProductByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (productUsecase *productUsecase) GetAllProduct(ctx context.Context) ([]product_entity.Product, error) {
	tx, err := productUsecase.db.Beginx()
	if err != nil {
		return nil, err
	}

	products, err := productUsecase.productRepository.GetAllProduct(ctx, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return products, nil
}

func (productUsecase *productUsecase) UpdateProductByID(ctx context.Context, id string, name string, price float64) error {
	tx, err := productUsecase.db.Beginx()
	if err != nil {
		return err
	}

	product, err := productUsecase.productRepository.GetProductByID(ctx, tx, id)
	if err != nil {
		return err
	}

	err = productUsecase.productRepository.UpdateProductByID(ctx, tx, product.ID, name, price)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (productUsecase *productUsecase) DeleteProductByID(ctx context.Context, id string) error {
	tx, err := productUsecase.db.Beginx()
	if err != nil {
		return err
	}

	product, err := productUsecase.productRepository.GetProductByID(ctx, tx, id)
	if err != nil {
		return err
	}

	err = productUsecase.productRepository.DeleteProductByID(ctx, tx, product.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
