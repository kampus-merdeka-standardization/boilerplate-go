package main

import (
	"log"

	"github.com/gin-gonic/gin"
	pinger_api "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/api"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
)

func main() {
	srv := httpPkg.NewHTTPServer(gin.DebugMode)
	root := srv.Group("")

	pinger_api.NewPingerController(root)

	err := srv.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
