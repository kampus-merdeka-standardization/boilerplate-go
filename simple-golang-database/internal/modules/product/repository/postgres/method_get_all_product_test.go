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

func TestGetAllProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed mocking sql")
	}
	defer db.Close()

	repoMock := NewProductRepository(sqlx.NewDb(db, "sqlmock"))

	expectedResults := []product_entity.Product{
		{
			ID:    "123",
			Name:  "abc",
			Price: 8000,
		},
	}

	t.Run("should return query result", func(t *testing.T) {
		sqlMock.ExpectQuery(getAllProduct).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow("123", "abc", 8000),
			)

		products, err := repoMock.GetAllProduct(context.Background())
		require.Nil(t, err)
		assert.Equal(t, expectedResults, products)
	})

	t.Run("should return error when failed insert query", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(getAllProduct).
			WillReturnError(rowErr)

		res, err := repoMock.GetAllProduct(context.Background())

		assert.Equal(t, []product_entity.Product(nil), res)
		assert.EqualError(t, err, rowErr.Error())
	})
}
