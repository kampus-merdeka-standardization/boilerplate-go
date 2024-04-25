package pkg_http_middleware

import (
	pkg_http "config-example/pkg/http"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIdAssignmentMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceContext := context.WithValue(ctx.Request.Context(), pkg_http.TraceString, uuid.NewString())
		httpReq := ctx.Request.WithContext(traceContext)
		ctx.Request = httpReq
	}
}
