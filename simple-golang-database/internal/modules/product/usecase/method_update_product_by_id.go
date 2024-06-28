package product_usecase

import (
	"context"
	"database/sql"
	pkg_error "simple-golang-database/pkg/error"
)

func (productUsecase *productUsecase) UpdateProductByID(ctx context.Context, id string, name string, price float64) error {
	product, err := productUsecase.productRepository.GetProductByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return pkg_error.NewNotFound(err, "Product Not found")
		}
		return err
	}

	err = productUsecase.productRepository.UpdateProductByID(ctx, product.ID, name, price)
	if err != nil {
		return err
	}

	return nil
}
