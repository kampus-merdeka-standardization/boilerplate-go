package pkg_http_middleware_test

import (
	pkg_http "config-example/pkg/http"
	pkg_http_middleware "config-example/pkg/http/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestTraceMiddleware(t *testing.T) {
	srv := pkg_http.NewHTTPServer(gin.TestMode)

	srv.Use(pkg_http_middleware.TraceIdAssignmentMiddleware())

	srv.GET("", func(ctx *gin.Context) {
		traceID, ok := ctx.Request.Context().Value(pkg_http.TraceString).(string)
		require.True(t, ok)
		require.NotEmpty(t, traceID)
	})

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	srv.ServeHTTP(res, req)
}
