package product_postgres

import (
	"context"
	"database/sql"
	product_model "simple-golang-database/internal/modules/product/model/entity"
)

func (productPostgresRepository *productPostgresRepository) DeleteProductByID(ctx context.Context, id string) error {
	tx, err := productPostgresRepository.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}
	defer tx.Commit()

	var product product_model.Product

	err = tx.GetContext(ctx, &product, getProductByID, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, deleteProductByID, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
