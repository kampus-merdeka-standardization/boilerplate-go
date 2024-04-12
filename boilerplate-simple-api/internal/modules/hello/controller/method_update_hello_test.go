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
	"github.com/stretchr/testify/require"
)

func TestUpdateHello(t *testing.T) {
	srv := setupTest()
	t.Run("Update Hello Http Test 1", func(t *testing.T) {
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
