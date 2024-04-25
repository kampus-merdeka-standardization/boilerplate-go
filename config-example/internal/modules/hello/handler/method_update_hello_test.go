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
	"github.com/stretchr/testify/require"
)

func TestUpdateHello(t *testing.T) {
	srv := setupTest()
	t.Run("Success - Should Return Expected Result", func(t *testing.T) {
		res := httptest.NewRecorder()

		reqBody := hello_request.UpdateHello{
			NewName: "Melza",
		}
		reqBytes, err := json.Marshal(reqBody)
		require.Nil(t, err)
		req := httptest.NewRequest(http.MethodPatch, "/hello", bytes.NewReader(reqBytes))

		srv.ServeHTTP(res, req)

		resBody := pkg_http_wrapper.NewResponse("")
		err = json.Unmarshal(res.Body.Bytes(), &resBody)
		require.Nil(t, err)

		messageExpect := fmt.Sprintf("Your name is replaced to %s", reqBody.NewName)
		assert.Equal(t, messageExpect, resBody.Message)
	})
}
