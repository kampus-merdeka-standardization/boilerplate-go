package product_postgres

import (
	"context"
	"database/sql"
	product_model "simple-golang-database/internal/modules/product/model/entity"
	errorPkg "simple-golang-database/pkg/error"
)

func (productPostgresRepository *productPostgresRepository) UpdateProductByID(ctx context.Context, id string, name string, price float64) error {
	tx, err := productPostgresRepository.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}

	var product product_model.Product

	err = tx.GetContext(ctx, &product, getProductByID, id)
	defer tx.Commit()

	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return errorPkg.NewNotFound(err, "Can't find product by that id")
		}
		return errorPkg.NewBadRequest(err, "Error While finding product by id")
	}

	_, err = tx.ExecContext(ctx, updateProductByID, id, name, price)
	if err != nil {
		return errorPkg.NewBadRequest(err, "Error while updating product by id")
	}

	return nil
}
