package main

import (
	pkg_http "simple-golang-monitoring/pkg/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := pkg_http.NewHTTPServer(gin.DebugMode)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
