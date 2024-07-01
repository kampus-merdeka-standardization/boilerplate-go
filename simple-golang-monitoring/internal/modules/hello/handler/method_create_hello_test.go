package hello_handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	hello_request "simple-golang-monitoring/internal/modules/hello/models/request"
	pkg_http_wrapper "simple-golang-monitoring/pkg/http/wrapper"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHello(t *testing.T) {
	srv := setupTest()

	t.Run("Success - Should Return Expected Result", func(t *testing.T) {
		res := httptest.NewRecorder()
		reqBody := hello_request.CreateHello{
			Name: "Azie",
			Age:  21,
		}
		reqBodyBytes, err := json.Marshal(reqBody)
		assert.Nil(t, err)
		req := httptest.NewRequest(http.MethodPost, "/hello", bytes.NewReader(reqBodyBytes))

		srv.ServeHTTP(res, req)

		assert.Equal(t, 200, res.Code)

		messageExpect := fmt.Sprintf("Hello, %s you are %d years old", reqBody.Name, reqBody.Age)

		resBody := pkg_http_wrapper.NewResponseWithValue(0, "", hello_request.CreateHello{})
		err = json.Unmarshal(res.Body.Bytes(), &resBody)
		assert.Nil(t, err)

		assert.Equal(t, resBody.Meta.Message, messageExpect)

		assert.Equal(t, reqBody, resBody.Value)
	})
}
