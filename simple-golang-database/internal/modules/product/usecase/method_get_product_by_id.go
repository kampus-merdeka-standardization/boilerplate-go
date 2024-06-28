package product_usecase

import (
	"context"
	"database/sql"
	product_entity "simple-golang-database/internal/modules/product/model/entity"
	pkg_error "simple-golang-database/pkg/error"
)

func (productUsecase *productUsecase) GetProductByID(ctx context.Context, id string) (*product_entity.Product, error) {
	product, err := productUsecase.productRepository.GetProductByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg_error.NewNotFound(err, "Product Not found")
		}
		return nil, err
	}

	return product, nil
}
