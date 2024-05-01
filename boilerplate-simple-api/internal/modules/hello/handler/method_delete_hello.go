package hello_handler

import (
	"fmt"
	"net/http"
	pkg_http_wrapper "simple-api/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// path : /hello/:id [DELETE]
func (hc *helloController) DeleteHello(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		pkg_http_wrapper.NewResponse(http.StatusOK,fmt.Sprintf("Your Data by the id of %s is successfully deleted", ctx.Param("id"))),
	)
}
