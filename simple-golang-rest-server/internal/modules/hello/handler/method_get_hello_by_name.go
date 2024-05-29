package hello_handler

import (
	"fmt"
	"net/http"
	pkg_http_wrapper "simple-api/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// path : /hello/:name [GET]
func (hc *helloController) GetHelloByName(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		pkg_http_wrapper.NewResponse(http.StatusOK,fmt.Sprintf("Hello, %s!", ctx.Param("name"))),
	)
}
