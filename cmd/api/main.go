package main

import (
	"github.com/gin-gonic/gin"

	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
)

func main() {
	srv := httpPkg.NewHTTPServer("debug")

	srv.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, httpPkg.Response{
			Message: "Pong!!!",
		})
	})

	srv.Run(":8080")
}
