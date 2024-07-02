package product_handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	mock_usecase "simple-golang-database/mock/usecase"
	pkg_http "simple-golang-database/pkg/http"
	pkg_http_middleware "simple-golang-database/pkg/http/middleware"
	pkg_http_wrapper "simple-golang-database/pkg/http/wrapper"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestHandlerDeleteProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mock_usecase.NewMockProductUsecase(ctrl)
	handlerMock := productController{
		productUsecase: usecaseMock,
	}

	url := "/product"
	urlWithParameter := url + "/:id"
	app := pkg_http.NewHTTPServer(gin.TestMode)
	app.Use(
		gin.Logger(),
		gin.Recovery(),
		pkg_http_middleware.ErrorHandler(),
		pkg_http_middleware.CorsHandler(),
	)
	app.DELETE(urlWithParameter, handlerMock.DeleteProductByID)

	paramID := "abc123-dfg456-jmq789-123bhd"

	t.Run("should return 200 with expected response", func(t *testing.T) {
		usecaseMock.EXPECT().DeleteProductByID(gomock.Any(), paramID).
			Return(nil)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", url, paramID), nil)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		res := pkg_http_wrapper.NewResponse(0, "")
		err := json.NewDecoder(w.Body).Decode(&res)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		require.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("should return 500 with expected response - error from usecase", func(t *testing.T) {
		expectedErr := errors.New("error from usecase")

		usecaseMock.EXPECT().
			DeleteProductByID(gomock.Any(), paramID).
			Return(expectedErr)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", url, paramID), nil)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		body := pkg_http_wrapper.NewResponse(0, "")
		_ = json.NewDecoder(w.Body).Decode(&body)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
