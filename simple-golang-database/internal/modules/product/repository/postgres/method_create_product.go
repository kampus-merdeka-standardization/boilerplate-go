package product_postgres

import (
	"context"

	"github.com/google/uuid"
)

func (productPostgresRepository *productPostgresRepository) CreateProduct(ctx context.Context, name string, price float64) (string, error) {
	id := uuid.NewString()

	_, err := productPostgresRepository.db.ExecContext(ctx, createProduct, id, name, price)
	if err != nil {
		return "", err
	}

	return id, nil
}
