package main

import (
	"github.com/gin-gonic/gin"
	pinger_graphql "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/graphql"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http/middleware"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
)

func main() {

	srv := httpPkg.NewHTTPServer(gin.DebugMode)
	log := logger.NewLogger(gin.DebugMode)

	srv.Use(middleware.LogHandler(log), gin.Recovery())
	srv.Use(middleware.CorsHandler())

	srv.POST("/pinger", pinger_graphql.NewPingerHandler)

	err := srv.Run(":8082")
	if err != nil {
		log.Fatal(err.Error())
	}
}
