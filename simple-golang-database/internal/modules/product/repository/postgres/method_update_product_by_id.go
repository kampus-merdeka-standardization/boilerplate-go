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

	var product product_entity.Product

	err = tx.GetContext(ctx, &product, getProductByID, id)
	defer tx.Commit()

	if err != nil {
		tx.Rollback()
		return product_entity.Product{}, err
	}

	row := tx.QueryRowxContext(ctx, updateProductByID, id, name, price)

	err = row.StructScan(&product)
	if err != nil {
		return product_entity.Product{}, err
	}

	return product, nil
}
