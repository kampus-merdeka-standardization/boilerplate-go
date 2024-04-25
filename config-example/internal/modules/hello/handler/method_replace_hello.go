package hello_handler

import (
	hello_request "config-example/internal/modules/hello/models/request"
	pkg_http_wrapper "config-example/pkg/http/wrapper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// path : /hello [PUT]
func (hc *helloController) ReplaceHello(ctx *gin.Context) {
	var req hello_request.ReplaceHello
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		pkg_http_wrapper.NewResponse(
			fmt.Sprintf("Your name is replaced from %s to %s", req.CurrentName, req.NewName),
		),
	)
}
