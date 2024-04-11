package hello_controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	pkg_http_wrapper "simple-api/pkg/http/wrapper"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteHello(t *testing.T) {
	srv := setupTest()

	t.Run("Delete Hello Http Test 1", func(t *testing.T) {
		res := httptest.NewRecorder()

		id := uuid.NewString()
		req := httptest.NewRequest(http.MethodDelete, "/hello/"+id, nil)

		srv.ServeHTTP(res, req)

		messageExpect := fmt.Sprintf("Your Data by the id of %s is successfully deleted", id)

		resBody := pkg_http_wrapper.NewResponse("")
		err := json.Unmarshal(res.Body.Bytes(), &resBody)
		assert.Nil(t, err)

		assert.Empty(t, resBody.Error)

		assert.Equal(t, messageExpect, resBody.Message)
	})
}
