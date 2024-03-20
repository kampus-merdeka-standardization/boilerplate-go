package main

import (
	"github.com/gin-gonic/gin"
	pinger_api "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/api"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/configs"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http/middleware"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
	"github.com/spf13/viper"
)

func main() {
	err := configs.SetEnvVariables("./configs/env/api.env")
	if err != nil {
		panic(err)
	}

	var conf configs.ApiConfig
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	srv := httpPkg.NewHTTPServer(conf.AppEnv)

	logger.InitLogger(conf.AppEnv, conf.LogPath)

	srv.Use(middleware.LogHandler(), gin.Recovery())
	srv.Use(middleware.CorsHandler())

	root := srv.Group("")

	pinger_api.NewPingerController(root)

	logger.Info("Running on Port " + conf.Port)
	if err := srv.Run(":" + conf.Port); err != nil {
		logger.Fatal(err.Error())
	}
}
