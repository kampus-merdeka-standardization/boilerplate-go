package hello_handler

import (
	hello_request "config-example/internal/modules/hello/models/request"
	pkg_http_wrapper "config-example/pkg/http/wrapper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// path : /hello [POST]
func (hc *helloController) CreateHello(ctx *gin.Context) {
	var req hello_request.CreateHello
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		pkg_http_wrapper.NewResponseWithValue(fmt.Sprintf("Hello, %s you are %d years old", req.Name, req.Age), req),
	)
}
