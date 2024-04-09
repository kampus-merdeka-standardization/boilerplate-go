package main

import (
	"fmt"
	hello_controller "simple-api/internal/modules/hello/controller"
	pkg_http "simple-api/pkg/http"

	"github.com/gin-gonic/gin"
)

const PORT int = 8000

func main() {
	srv := pkg_http.NewHTTPServer(gin.DebugMode)

	hello_controller.BindHelloController(srv.Group("/hello"))

	srv.Run(fmt.Sprintf(":%d", PORT))
}
