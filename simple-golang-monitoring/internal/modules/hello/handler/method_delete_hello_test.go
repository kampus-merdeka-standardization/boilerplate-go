package hello_handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	pkg_http_wrapper "simple-golang-monitoring/pkg/http/wrapper"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteHello(t *testing.T) {
	srv := setupTest()

	t.Run("Success - Should Return Expected Result", func(t *testing.T) {
		res := httptest.NewRecorder()

		id := uuid.NewString()
		req := httptest.NewRequest(http.MethodDelete, "/hello/"+id, nil)

		srv.ServeHTTP(res, req)

		messageExpect := fmt.Sprintf("Your Data by the id of %s is successfully deleted", id)

		resBody := pkg_http_wrapper.NewResponse(0, "")
		err := json.Unmarshal(res.Body.Bytes(), &resBody)
		assert.Nil(t, err)

		assert.Equal(t, messageExpect, resBody.Meta.Message)
		assert.Equal(t, http.StatusOK, resBody.Meta.Code)
	})
}
