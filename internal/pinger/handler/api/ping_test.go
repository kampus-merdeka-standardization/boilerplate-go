package pinger_api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	pinger_api "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/api"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
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
