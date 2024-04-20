package hello_handler

import (
	"fmt"
	"net/http"
	hello_request "simple-api/internal/modules/hello/models/request"
	pkg_http_wrapper "simple-api/pkg/http/wrapper"

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
