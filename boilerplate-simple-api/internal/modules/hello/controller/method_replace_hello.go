package hello_controller

import (
	"fmt"
	"net/http"
	hello_request "simple-api/internal/modules/hello/models/request"
	pkg_http_wrapper "simple-api/pkg/http/wrapper"

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
