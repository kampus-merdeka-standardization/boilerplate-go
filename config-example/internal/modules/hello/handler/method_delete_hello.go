package hello_handler

import (
	pkg_http_wrapper "config-example/pkg/http/wrapper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// path : /hello/:id [DELETE]
func (hc *helloController) DeleteHello(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		pkg_http_wrapper.NewResponse(fmt.Sprintf("Your Data by the id of %s is successfully deleted", ctx.Param("id"))),
	)
}
