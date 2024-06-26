package hello_handler

import (
	"fmt"
	"net/http"
	hello_request "simple-golang-monitoring/internal/modules/hello/models/request"
	pkg_http_wrapper "simple-golang-monitoring/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// path : /hello [PUT]
func (hc *helloController) ReplaceHello(ctx *gin.Context) {
	_, span := hc.tracer.StartTransaction(ctx.Request.Context(), "Replace Hello Handler")
	defer hc.tracer.EndTransaction(span)

	var req hello_request.ReplaceHello
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		pkg_http_wrapper.NewResponse(
			http.StatusOK,
			fmt.Sprintf("Your name is replaced from %s to %s", req.CurrentName, req.NewName),
		),
	)
}
