package product_postgres

import (
	"context"
	"database/sql"
	product_entity "simple-golang-database/internal/modules/product/model/entity"
)

func (productPostgresRepository *productPostgresRepository) UpdateProductByID(ctx context.Context, id string, name string, price int64) (product_entity.Product, error) {
	tx, err := productPostgresRepository.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
	if err != nil {
		return product_entity.Product{}, nil
	}
	defer tx.Commit()

	var product product_entity.Product

	row := productPostgresRepository.db.QueryRowContext(ctx, getProductByID, id)

	err = row.Scan(&product)

	if err != nil {
		tx.Rollback()
		return product_entity.Product{}, err
	}

	err = tx.GetContext(ctx, &product, updateProductByID, id, name, price)
	if err != nil {
		return product_entity.Product{}, err
	}

	return product, nil
}
