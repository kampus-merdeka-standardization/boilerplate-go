package product_postgres

import (
	"context"
	product_model "simple-golang-database/internal/modules/product/model/entity"
)

func (productPostgresRepository *productPostgresRepository) GetAllProduct(ctx context.Context) ([]product_model.Product, error) {
	var products []product_model.Product

	err := productPostgresRepository.db.SelectContext(ctx, &products, getAllProduct)
	if err != nil {
		return nil, err
	}

	return products, nil
}
