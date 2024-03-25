package product_postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	product_model "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/model/entity"
	errorPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/error"
)

func (productPostgresRepository *productPostgresRepository) CreateProduct(ctx context.Context, tx *sqlx.Tx, name string, price float64) (string, error) {
	id := uuid.NewString()
	_, err := tx.ExecContext(ctx, createProduct, id, name, price)
	if err != nil {
		return "", errorPkg.NewBadRequest(err, "Error while creating product")
	}

	return id, nil
}

func (productPostgresRepository *productPostgresRepository) GetProductByID(ctx context.Context, tx *sqlx.Tx, id string) (*product_model.Product, error) {
	var product product_model.Product

	err := tx.GetContext(ctx, &product, getProductByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorPkg.NewNotFound(err, "Product Not Found")
		}
		return nil, errorPkg.NewBadRequest(err, "Error while getting product by id")
	}

	return &product, nil
}

func (productPostgresRepository *productPostgresRepository) GetAllProduct(ctx context.Context, tx *sqlx.Tx) ([]product_model.Product, error) {
	var products []product_model.Product

	err := tx.SelectContext(ctx, &products, getAllProduct)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorPkg.NewNotFound(err, "Product Not Found")
		}
		return nil, errorPkg.NewBadRequest(err, "Error while getting all product")
	}

	return products, nil
}

func (productPostgresRepository *productPostgresRepository) UpdateProductByID(ctx context.Context, tx *sqlx.Tx, id string, name string, price float64) error {
	_, err := tx.ExecContext(ctx, updateProductByID, id, name, price)
	if err != nil {
		return errorPkg.NewBadRequest(err, "Error while updating product by id")
	}

	return nil
}

func (productPostgresRepository *productPostgresRepository) DeleteProductByID(ctx context.Context, tx *sqlx.Tx, id string) error {
	_, err := tx.ExecContext(ctx, deleteProductByID, id)
	if err != nil {
		return errorPkg.NewBadRequest(err, "Error while deleting product by id")
	}

	return nil
}
