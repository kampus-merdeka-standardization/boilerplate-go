package product_postgres

import (
	"context"
	"errors"
	product_entity "simple-golang-database/internal/modules/product/model/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed mocking sql")
	}
	defer db.Close()

	repoMock := NewProductRepository(sqlx.NewDb(db, "sqlmock"))

	expectedResult := product_entity.Product{
		ID:    "123",
		Name:  "Tea",
		Price: 25000,
	}

	t.Run("should return query result", func(t *testing.T) {
		sqlMock.ExpectQuery(getProductByID).WithArgs(expectedResult.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow("123", "Tea", 25000),
			)

		product, err := repoMock.GetProductByID(context.Background(), expectedResult.ID)
		require.Nil(t, err)
		assert.Equal(t, expectedResult, product)
	})

	t.Run("should return error when failed insert query", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(getProductByID).
			WillReturnError(rowErr)

		res, err := repoMock.GetProductByID(context.Background(), "")

		assert.Equal(t, product_entity.Product{}, res)
		assert.EqualError(t, err, rowErr.Error())
	})

	t.Run("should return error when get last inserted id", func(t *testing.T) {
		rowErr := errors.New("error from db")

		sqlMock.ExpectQuery(getProductByID).
			WithArgs(expectedResult.ID).
			WillReturnError(rowErr)

		res, err := repoMock.GetProductByID(context.Background(), expectedResult.ID)

		assert.Equal(t, product_entity.Product{}, res)
		assert.EqualError(t, err, rowErr.Error())
	})
}
