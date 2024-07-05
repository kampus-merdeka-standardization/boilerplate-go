package product_handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	product_response "simple-golang-database/internal/modules/product/model/response"
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

func TestHandlerGetProductByID(t *testing.T) {
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
	app.POST(urlWithParameter, handlerMock.GetProductByID)

	paramID := "abc123-dfg456-jmq789-123bhd"

	expectedRes := product_response.Product{
		ID:    "abc1d-1lakd1",
		Name:  "Handuk",
		Price: 125000,
	}

	t.Run("should return 200 with expected response", func(t *testing.T) {
		usecaseMock.EXPECT().GetProductByID(gomock.Any(), paramID).
			Return(expectedRes, nil)

		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url, paramID), nil)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		res := pkg_http_wrapper.NewResponseWithValue(0, "", product_response.Product{})
		err := json.NewDecoder(w.Body).Decode(&res)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		require.Nil(t, err)
		assert.Equal(t, expectedRes, res.Value)
	})

	t.Run("should return 500 with expected response - error from usecase", func(t *testing.T) {
		expectedErr := errors.New("error from usecase")

		usecaseMock.EXPECT().
			GetProductByID(gomock.Any(), paramID).
			Return(product_response.Product{}, expectedErr)

		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url, paramID), nil)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		body := pkg_http_wrapper.NewResponseWithValue(0, "", product_response.Product{})
		_ = json.NewDecoder(w.Body).Decode(&body)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
