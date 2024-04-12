package hello_controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	pkg_http_wrapper "simple-api/pkg/http/wrapper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHelloByName(t *testing.T) {
	srv := setupTest()

	t.Run("Get Hello By Name Http Test 1", func(t *testing.T) {
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/hello/azie", nil)

		srv.ServeHTTP(res, req)

		name := "azie"
		messageExpect := fmt.Sprintf("Hello, %s!", name)

		resBody := pkg_http_wrapper.NewResponse("")
		err := json.Unmarshal(res.Body.Bytes(), &resBody)
		assert.Nil(t, err)

		assert.Equal(t, messageExpect, resBody.Message)
	})
}
