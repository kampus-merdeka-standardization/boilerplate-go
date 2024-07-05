package product_handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	product_request "simple-golang-database/internal/modules/product/model/request"
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

func TestHandlerCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mock_usecase.NewMockProductUsecase(ctrl)
	handlerMock := productController{
		productUsecase: usecaseMock,
	}

	url := "/product"
	app := pkg_http.NewHTTPServer(gin.TestMode)
	app.Use(
		gin.Logger(),
		gin.Recovery(),
		pkg_http_middleware.ErrorHandler(),
		pkg_http_middleware.CorsHandler(),
	)
	app.POST(url, handlerMock.CreateProduct)

	param := product_request.CreateProduct{
		Name:  "Handuk",
		Price: 125000,
	}
	bodyRequest, _ := json.Marshal(param)

	expectedRes := product_response.Product{
		ID:    "abc1d-1lakd1",
		Name:  "Handuk",
		Price: 125000,
	}

	t.Run("should return 200 with expected response", func(t *testing.T) {
		usecaseMock.EXPECT().CreateProduct(gomock.Any(), param.Name, param.Price).Return(expectedRes, nil)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		res := pkg_http_wrapper.NewResponseWithValue(0, "", product_response.Product{})
		err := json.NewDecoder(w.Body).Decode(&res)

		assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
		require.Nil(t, err)
		assert.Equal(t, expectedRes, res.Value)
	})

	t.Run("should return 400 - invalid body request", func(t *testing.T) {
		bodyRequest, _ := json.Marshal(map[string]string{})

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		res := pkg_http_wrapper.NewResponseWithValue(0, "", product_response.Product{})
		err := json.NewDecoder(w.Body).Decode(&res)
		require.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 500 with expected response - error from usecase", func(t *testing.T) {
		expectedErr := errors.New("error from usecase")

		usecaseMock.EXPECT().
			CreateProduct(gomock.Any(), param.Name, param.Price).
			Return(product_response.Product{}, expectedErr)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		body := pkg_http_wrapper.NewResponse(0, "")
		_ = json.NewDecoder(w.Body).Decode(&body)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
