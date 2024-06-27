package main

import (
	"fmt"
	hello_handler "simple-golang-rest-server/internal/modules/hello/handler"
	pkg_http "simple-golang-rest-server/pkg/http"

	"github.com/gin-gonic/gin"
)

const PORT int = 8000

func main() {
	srv := pkg_http.NewHTTPServer(gin.DebugMode)

	hello_handler.BindHelloHandler(srv.Group("/hello"))

	srv.Run(fmt.Sprintf(":%d", PORT))
}
