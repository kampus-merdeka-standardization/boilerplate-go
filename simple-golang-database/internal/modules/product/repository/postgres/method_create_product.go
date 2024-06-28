package product_postgres

import (
	"context"
	product_entity "simple-golang-database/internal/modules/product/model/entity"
)

func (productPostgresRepository *productPostgresRepository) CreateProduct(ctx context.Context, name string, price int64) (product_entity.Product, error) {
	var product product_entity.Product

	err := productPostgresRepository.db.GetContext(ctx, &product, createProduct, name, price)
	if err != nil {
		return product_entity.Product{}, err
	}

	return product, nil
}
