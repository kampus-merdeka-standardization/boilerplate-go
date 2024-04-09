package main

import (
	"fmt"
	pkg_http "simple-api/pkg/http"

	"github.com/gin-gonic/gin"
)

const PORT int = 8000

func main() {
	srv := pkg_http.NewHTTPServer(gin.DebugMode)

	srv.Run(fmt.Sprintf(":%d", PORT))
}
