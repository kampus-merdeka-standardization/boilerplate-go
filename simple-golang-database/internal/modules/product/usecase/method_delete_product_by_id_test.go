package product_usecase

import (
	"context"
	"errors"
	mock_repository "simple-golang-database/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestDeleteProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockProductRepository(ctrl)
	usecaseMock := productUsecase{
		productRepository: repoMock,
	}

	idParam := "c5f08a52-cc46-47d1-879b-15e120885366"

	t.Run("should delete product", func(t *testing.T) {
		repoMock.EXPECT().DeleteProductByID(gomock.Any(), idParam).
			Return(nil)

		err := usecaseMock.DeleteProductByID(context.Background(), idParam)
		require.Nil(t, err)
	})

	t.Run("should return error when create new product is failed", func(t *testing.T) {
		expectedErr := errors.New("db down")

		repoMock.EXPECT().DeleteProductByID(gomock.Any(), idParam).
			Return(expectedErr)

		err := usecaseMock.DeleteProductByID(context.Background(), idParam)

		assert.Equal(t, expectedErr, err)
	})
}
