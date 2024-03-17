package main

import (
	"github.com/gin-gonic/gin"
	pinger_api "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/api"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http/middleware"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
)

func main() {
	srv := httpPkg.NewHTTPServer(gin.DebugMode)

	log := logger.NewLogger(gin.DebugMode)

	srv.Use(middleware.LogHandler(log), gin.Recovery())

	root := srv.Group("")

	pinger_api.NewPingerController(root)

	err := srv.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
