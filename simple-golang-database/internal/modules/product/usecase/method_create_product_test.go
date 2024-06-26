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

func TestUsecaseCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockProductRepository(ctrl)
	usecaseMock := productUsecase{
		productRepository: repoMock,
	}

	nameParam := "Laptop Lenovo XYZ"
	priceParam := int64(30000000)

	expectedRepoReturn := product_entity.Product{
		ID:    "c5f08a52-cc46-47d1-879b-15e120885366",
		Name:  nameParam,
		Price: priceParam,
	}

	t.Run("should create new product", func(t *testing.T) {
		repoMock.EXPECT().CreateProduct(gomock.Any(), nameParam, priceParam).
			Return(expectedRepoReturn, nil)

		res, err := usecaseMock.CreateProduct(context.Background(), nameParam, priceParam)

		require.Nil(t, err)
		assert.Equal(t, expectedRepoReturn.ID, res.ID)
		assert.Equal(t, expectedRepoReturn.Name, res.Name)
		assert.Equal(t, expectedRepoReturn.Price, res.Price)
	})

	t.Run("should return error when create new product is failed", func(t *testing.T) {
		expectedErr := errors.New("db down")

		repoMock.EXPECT().CreateProduct(gomock.Any(), nameParam, priceParam).
			Return(product_entity.Product{}, expectedErr)

		res, err := usecaseMock.CreateProduct(context.Background(), nameParam, priceParam)

		assert.Equal(t, expectedErr, err)
		assert.Equal(t, "", res.ID)
	})
}
