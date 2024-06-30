package product_usecase

import (
	"context"
	product_entity "simple-golang-database/internal/modules/product/model/entity"
	mock_repository "simple-golang-database/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockProductRepository(ctrl)
	usecaseMock := productUsecase{
		productRepository: repoMock,
	}

	idParam := "c5f08a52-cc46-47d1-879b-15e120885366"

	expectedRepoReturn := product_entity.Product{
		ID:    idParam,
		Name:  "Biore Handwash",
		Price: 35000,
	}

	t.Run("should return product by id", func(t *testing.T) {
		repoMock.EXPECT().GetProductByID(gomock.Any(), idParam).
			Return(expectedRepoReturn, nil)

		res, err := usecaseMock.GetProductByID(context.Background(), idParam)
		require.Nil(t, err)
		assert.Equal(t, expectedRepoReturn.ID, res.ID)
		assert.Equal(t, expectedRepoReturn.Name, res.Name)
		assert.Equal(t, expectedRepoReturn.Price, res.Price)
	})
}
