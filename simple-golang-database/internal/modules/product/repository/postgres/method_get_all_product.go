package product_postgres

import (
	"context"
	"database/sql"
	product_model "simple-golang-database/internal/modules/product/model/entity"
	errorPkg "simple-golang-database/pkg/error"
)

func (productPostgresRepository *productPostgresRepository) GetAllProduct(ctx context.Context) ([]product_model.Product, error) {
	var products []product_model.Product

	err := productPostgresRepository.db.SelectContext(ctx, &products, getAllProduct)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorPkg.NewNotFound(err, "Product Not Found")
		}
		return nil, errorPkg.NewBadRequest(err, "Error while getting all product")
	}

	return products, nil
}
