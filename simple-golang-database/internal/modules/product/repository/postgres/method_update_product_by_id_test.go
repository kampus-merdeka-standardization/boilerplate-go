package product_postgres

import (
	"context"
	"database/sql"
	"errors"
	product_entity "simple-golang-database/internal/modules/product/model/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRepositoryUpdateProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed mocking sql")
	}
	defer db.Close()

	repoMock := NewProductRepository(sqlx.NewDb(db, "sqlmock"))

	oldRow := product_entity.Product{
		ID:    "123",
		Name:  "Coffee",
		Price: 50000,
	}

	expecteRow := product_entity.Product{
		ID:    "123",
		Name:  "Tea",
		Price: 25000,
	}

	t.Run("should execute update query", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectQuery(getProductByID).
			WithArgs(oldRow.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(oldRow.ID, oldRow.Name, oldRow.Price),
			)

		sqlMock.ExpectQuery(updateProductByID).
			WithArgs(expecteRow.ID, expecteRow.Name, expecteRow.Price).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(expecteRow.ID, expecteRow.Name, expecteRow.Price),
			)

		sqlMock.ExpectCommit()

		product, err := repoMock.UpdateProductByID(context.Background(), expecteRow.ID, expecteRow.Name, expecteRow.Price)
		require.Nil(t, err)
		assert.Equal(t, expecteRow, product)
	})

	t.Run("should return error when failed update query", func(t *testing.T) {
		dbErr := errors.New("database error")
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(getProductByID).
			WithArgs(expecteRow.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(oldRow.ID, oldRow.Name, oldRow.Price),
			)
		sqlMock.ExpectQuery(updateProductByID).WithArgs(expecteRow.ID, expecteRow.Name, expecteRow.Price).
			WillReturnError(dbErr)
		sqlMock.ExpectRollback()

		_, err := repoMock.UpdateProductByID(context.Background(), expecteRow.ID, expecteRow.Name, expecteRow.Price)
		assert.EqualError(t, err, dbErr.Error())
	})

	t.Run("should return error when id is not found", func(t *testing.T) {
		noRowsErr := sql.ErrNoRows
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(getProductByID).
			WithArgs(expecteRow.ID).
			WillReturnError(noRowsErr)

		sqlMock.ExpectRollback()

		_, err := repoMock.UpdateProductByID(context.Background(), expecteRow.ID, expecteRow.Name, expecteRow.Price)
		assert.Equal(t, err, err)
	})
}
