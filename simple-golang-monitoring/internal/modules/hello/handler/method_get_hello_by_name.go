package hello_handler

import (
	"fmt"
	"net/http"
	pkg_http_wrapper "simple-golang-monitoring/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// path : /hello/:name [GET]
func (hc *helloController) GetHelloByName(c *gin.Context) {
	_, span := hc.tracer.StartTransaction(c.Request.Context(), "Get Hello By Name Handler")
	defer hc.tracer.EndTransaction(span)

	c.JSON(
		http.StatusOK,
		pkg_http_wrapper.NewResponse(http.StatusOK, fmt.Sprintf("Hello, %s!", c.Param("name"))),
	)
}
