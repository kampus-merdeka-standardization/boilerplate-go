package pinger_api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	pinger_api "simple-golang-app/internal/modules/pinger/handler/api"
	httpPkg "simple-golang-app/pkg/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTest() *gin.Engine {
	router := httpPkg.NewHTTPServer("test")
	pinger_api.NewPingerController(router.Group("/ping"))

	return router
}

func TestPingApi(t *testing.T) {
	router := setupTest()

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)

	var resBody httpPkg.Response
	json.Unmarshal(res.Body.Bytes(), &resBody)
	assert.Equal(t, httpPkg.Response{
		Message: "pong",
	}, resBody)
}
