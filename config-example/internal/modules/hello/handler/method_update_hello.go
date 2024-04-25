package hello_handler

import (
	hello_request "config-example/internal/modules/hello/models/request"
	pkg_http_wrapper "config-example/pkg/http/wrapper"
	"fmt"
	"net/http"

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
		pkg_http_wrapper.NewResponse(fmt.Sprintf("Your name is replaced to %s", req.NewName)),
	)
}
