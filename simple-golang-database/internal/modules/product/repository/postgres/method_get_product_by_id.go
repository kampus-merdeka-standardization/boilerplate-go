package product_postgres

import (
	"context"
	"database/sql"
	product_model "simple-golang-database/internal/modules/product/model/entity"
	errorPkg "simple-golang-database/pkg/error"
)

func (productPostgresRepository *productPostgresRepository) GetProductByID(ctx context.Context, id string) (*product_model.Product, error) {
	var product product_model.Product

	err := productPostgresRepository.db.GetContext(ctx, &product, getProductByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorPkg.NewNotFound(err, "Product Not Found")
		}
		return nil, errorPkg.NewBadRequest(err, "Error while getting product by id")
	}

	return &product, nil
}
