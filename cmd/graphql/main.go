package main

import (
	"github.com/gin-gonic/gin"
	pinger_graphql "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/graphql"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/configs"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http/middleware"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
)

func main() {
	conf := configs.LoadGraphqlConfig()

	srv := httpPkg.NewHTTPServer(conf.AppEnv)

	logger.InitLogger(conf.AppEnv, conf.LogPath)

	srv.Use(middleware.LogHandler(), gin.Recovery())
	srv.Use(middleware.CorsHandler())

	srv.POST("/pinger", pinger_graphql.NewPingerHandler)

	logger.Info("Running on Port " + conf.Port)
	if err := srv.Run(":" + conf.Port); err != nil {
		logger.Fatal(err.Error())
	}
}
