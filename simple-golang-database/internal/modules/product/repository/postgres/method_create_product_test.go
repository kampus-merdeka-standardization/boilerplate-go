package product_postgres

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRepositoryCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed mocking sql")
	}
	defer db.Close()

	repoMock := NewProductRepository(sqlx.NewDb(db, "sqlmock"))

	reqName := "Tea Cup"
	reqPrice := float64(30000)

	t.Run("should execute insert query", func(t *testing.T) {
		sqlMock.ExpectExec(createProduct).WithArgs("", reqName, reqPrice)

		id, err := repoMock.CreateProduct(context.Background(), reqName, reqPrice)

		require.NotNil(t, err)
		assert.NotNil(t, id)
	})
}
