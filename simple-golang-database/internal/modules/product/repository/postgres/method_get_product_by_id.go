package product_postgres

import (
	"context"
	product_model "simple-golang-database/internal/modules/product/model/entity"
)

func (productPostgresRepository *productPostgresRepository) GetProductByID(ctx context.Context, id string) (product_model.Product, error) {
	var product product_model.Product

	row := productPostgresRepository.db.QueryRowContext(ctx, getProductByID, id)

	err := row.Scan(&product)
	if err != nil {
		return product_model.Product{}, err
	}

	return product, nil
}
