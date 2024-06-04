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

		resBody := pkg_http_wrapper.NewResponse(0, "")
		err = json.Unmarshal(res.Body.Bytes(), &resBody)
		assert.Nil(t, err)

		messageExpect := fmt.Sprintf("Your name is replaced from %s to %s", reqBody.CurrentName, reqBody.NewName)
		assert.Equal(t, messageExpect, resBody.Meta.Message)
		assert.Equal(t, http.StatusOK, resBody.Meta.Code)
	})
}
