package hello_handler

import (
	"fmt"
	"net/http"
	hello_request "simple-golang-monitoring/internal/modules/hello/models/request"
	pkg_http_wrapper "simple-golang-monitoring/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// path : /hello [PATCH]
func (hc *helloController) UpdateHello(ctx *gin.Context) {
	var req hello_request.UpdateHello
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		pkg_http_wrapper.NewResponse(http.StatusOK, fmt.Sprintf("Your name is replaced to %s", req.NewName)),
	)
}
