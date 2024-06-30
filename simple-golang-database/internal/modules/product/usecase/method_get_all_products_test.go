package product_usecase

import (
	"context"
	"errors"
	product_entity "simple-golang-database/internal/modules/product/model/entity"
	mock_repository "simple-golang-database/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAllProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockProductRepository(ctrl)
	usecaseMock := productUsecase{
		productRepository: repoMock,
	}

	expectedResults := []product_entity.Product{
		{
			ID:    "123",
			Name:  "abc",
			Price: 8000,
		},
	}

	t.Run("should get all products", func(t *testing.T) {
		repoMock.EXPECT().GetAllProduct(gomock.Any()).
			Return(expectedResults, nil)

		products, err := usecaseMock.GetAllProduct(context.Background())
		require.Nil(t, err)
		assert.True(t, len(products) > 0)
	})

	t.Run("should return error when create new product is failed", func(t *testing.T) {
		expectedErr := errors.New("db down")

		repoMock.EXPECT().GetAllProduct(gomock.Any()).
			Return(nil, expectedErr)

		res, err := usecaseMock.GetAllProduct(context.Background())

		assert.Equal(t, expectedErr, err)
		assert.Nil(t, res)
	})
}
