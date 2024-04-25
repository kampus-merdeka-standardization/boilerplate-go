package hello_handler_test

import (
	"bytes"
	hello_request "config-example/internal/modules/hello/models/request"
	pkg_http_wrapper "config-example/pkg/http/wrapper"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceHello(t *testing.T) {
	srv := setupTest()

	t.Run("Success - Should Return Expected Result", func(t *testing.T) {
		res := httptest.NewRecorder()

		reqBody := hello_request.ReplaceHello{
			CurrentName: "Azie",
			NewName:     "Melza",
		}
		reqBodyBytes, err := json.Marshal(reqBody)
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPut, "/hello", bytes.NewReader(reqBodyBytes))

		srv.ServeHTTP(res, req)

		resBody := pkg_http_wrapper.NewResponse("")
		err = json.Unmarshal(res.Body.Bytes(), &resBody)
		assert.Nil(t, err)

		messageExpect := fmt.Sprintf("Your name is replaced from %s to %s", reqBody.CurrentName, reqBody.NewName)
		assert.Equal(t, messageExpect, resBody.Message)
	})
}
