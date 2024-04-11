package hello_controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	hello_request "simple-api/internal/modules/hello/models/request"
	pkg_http_wrapper "simple-api/pkg/http/wrapper"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHello(t *testing.T) {
	srv := setupTest()

	t.Run("Create Hello Http Test 1", func(t *testing.T) {
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

		resBody := pkg_http_wrapper.NewResponse("", nil)
		err = json.Unmarshal(res.Body.Bytes(), &resBody)
		assert.Nil(t, err)

		assert.Equal(t, resBody.Message, messageExpect)
	})
}
