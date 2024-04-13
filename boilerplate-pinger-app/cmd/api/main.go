package main

import (
	"github.com/gin-gonic/gin"
	pinger_api "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/modules/pinger/handler/api"
	"github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/pkg/configs"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/pkg/http"
	"github.com/kampus-merdeka-standardization/boilerplate-pinger-app/pkg/http/middleware"
	"github.com/kampus-merdeka-standardization/boilerplate-pinger-app/pkg/logger"
)

func main() {
	conf := configs.LoadApiConfig()

	srv := httpPkg.NewHTTPServer(conf.AppEnv)

	logger.InitLogger(conf.AppEnv, conf.LogPath)

	srv.Use(middleware.LogHandler(), gin.Recovery())
	srv.Use(middleware.CorsHandler())
	srv.Use(middleware.ErrorHandler())

	router := srv.Group("")

	// pgDb := db.NewPostgresDB(db.PostgresDsn{
	// 	Host:     conf.PostgresHost,
	// 	User:     conf.PostgresUser,
	// 	Password: conf.PostgresPassword,
	// 	Port:     conf.PostgresPort,
	// 	Db:       conf.PostgresDb,
	// })

	pinger_api.NewPingerController(router.Group("/ping"))

	logger.Info("Running on Port " + conf.Port)
	if err := srv.RunTLS(":"+conf.Port, conf.CertFilePath, conf.KeyFilePath); err != nil {
		logger.Fatal(err.Error())
	}
}
