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

func TestRepositoryDeleteProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed mocking sql")
	}
	defer db.Close()

	repoMock := NewProductRepository(sqlx.NewDb(db, "sqlmock"))

	expecteRow := product_entity.Product{
		ID:    "123",
		Name:  "Tea",
		Price: 25000,
	}

	t.Run("should execute insert query", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(getProductByID).
			WithArgs(expecteRow.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(expecteRow.ID, expecteRow.Name, expecteRow.Price),
			)
		sqlMock.ExpectExec(deleteProductByID).WithArgs(expecteRow.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		sqlMock.ExpectCommit()

		err := repoMock.DeleteProductByID(context.Background(), expecteRow.ID)
		require.Nil(t, err)
	})

	t.Run("should return error when failed delete query", func(t *testing.T) {
		dbErr := errors.New("database error")
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(getProductByID).
			WithArgs(expecteRow.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(expecteRow.ID, expecteRow.Name, expecteRow.Price),
			)
		sqlMock.ExpectExec(deleteProductByID).
			WillReturnError(dbErr)
		sqlMock.ExpectRollback()

		err := repoMock.DeleteProductByID(context.Background(), expecteRow.ID)
		assert.EqualError(t, err, dbErr.Error())
	})

	t.Run("should return error when id is not found", func(t *testing.T) {
		noRowsErr := sql.ErrNoRows
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(getProductByID).
			WithArgs(expecteRow.ID).
			WillReturnError(noRowsErr)

		sqlMock.ExpectRollback()

		err := repoMock.DeleteProductByID(context.Background(), expecteRow.ID)
		assert.EqualError(t, err, noRowsErr.Error())
	})
}
