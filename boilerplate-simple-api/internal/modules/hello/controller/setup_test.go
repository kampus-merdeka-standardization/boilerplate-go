package hello_controller_test

import (
	hello_controller "simple-api/internal/modules/hello/controller"
	pkg_http "simple-api/pkg/http"

	"github.com/gin-gonic/gin"
)

func setupTest() *gin.Engine {
	srv := pkg_http.NewHTTPServer(gin.TestMode)

	hello_controller.BindHelloController(srv.Group("/hello"))
	return srv
}
