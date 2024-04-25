package hello_handler

import (
	hello_response "config-example/internal/modules/hello/models/response"
	pkg_http_wrapper "config-example/pkg/http/wrapper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// path : /hello/:name [GET]
func (hc *helloController) GetHelloByName(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(
		http.StatusOK,
		pkg_http_wrapper.NewResponseWithValue(fmt.Sprintf("Hello, %s!", name), hello_response.GetHelloByName{
			Name:  name,
			Hello: fmt.Sprintf("Hello to %s!!!", name),
		}),
	)
}
