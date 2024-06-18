package product_postgres

import (
	"context"

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
