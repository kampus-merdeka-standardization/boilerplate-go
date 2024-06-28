package product_usecase

import (
	"context"
	"database/sql"
	pkg_error "simple-golang-database/pkg/error"
)

func (productUsecase *productUsecase) DeleteProductByID(ctx context.Context, id string) error {
	product, err := productUsecase.productRepository.GetProductByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return pkg_error.NewNotFound(err, "Product Not found")
		}
		return err
	}

	err = productUsecase.productRepository.DeleteProductByID(ctx, product.ID)
	if err != nil {
		return err
	}

	return nil
}
