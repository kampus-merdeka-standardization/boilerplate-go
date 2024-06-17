package product_postgres

import (
	"context"
	"database/sql"

	product_model "simple-golang-database/internal/modules/product/model/entity"
	errorPkg "simple-golang-database/pkg/error"

	"github.com/google/uuid"
)

func (productPostgresRepository *productPostgresRepository) CreateProduct(ctx context.Context, name string, price float64) (string, error) {
	id := uuid.NewString()

	_, err := productPostgresRepository.db.ExecContext(ctx, createProduct, id, name, price)
	if err != nil {
		return "", errorPkg.NewBadRequest(err, "Error while creating product")
	}

	return id, nil
}

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

func (productPostgresRepository *productPostgresRepository) DeleteProductByID(ctx context.Context, id string) error {
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

	_, err = tx.ExecContext(ctx, deleteProductByID, id)
	if err != nil {
		return errorPkg.NewBadRequest(err, "Error while deleting product by id")
	}

	return nil
}
