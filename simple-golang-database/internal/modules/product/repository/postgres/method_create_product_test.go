package product_postgres

import (
	"context"
	"errors"
	product_entity "simple-golang-database/internal/modules/product/model/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
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

	expectedResult := product_entity.Product{
		ID:    uuid.NewString(),
		Name:  "Tea",
		Price: 25000,
	}

	t.Run("should execute insert query", func(t *testing.T) {
		sqlMock.ExpectQuery(createProduct).
			WithArgs(expectedResult.Name, expectedResult.Price).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow(expectedResult.ID, expectedResult.Name, expectedResult.Price))

		product, err := repoMock.CreateProduct(context.Background(), expectedResult.Name, expectedResult.Price)

		require.Nil(t, err)

		assert.Equal(t, expectedResult, product)
	})

	t.Run("should return error when failed insert query", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(createProduct).
			WillReturnError(rowErr)

		res, err := repoMock.CreateProduct(context.Background(), "", 0)

		assert.Equal(t, product_entity.Product{}, res)
		assert.EqualError(t, err, rowErr.Error())
	})

	t.Run("should return error when get last inserted id", func(t *testing.T) {
		rowErr := errors.New("error from db")

		sqlMock.ExpectQuery(createProduct).
			WithArgs(expectedResult.Name, expectedResult.Price).
			WillReturnError(rowErr)

		res, err := repoMock.CreateProduct(context.Background(), expectedResult.Name, expectedResult.Price)

		assert.Equal(t, product_entity.Product{}, res)
		assert.EqualError(t, err, rowErr.Error())
	})
}
