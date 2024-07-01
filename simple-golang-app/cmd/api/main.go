package main

import (
	pinger_api "simple-golang-app/internal/modules/pinger/handler/api"
	"simple-golang-app/internal/pkg/configs"
	httpPkg "simple-golang-app/pkg/http"
	"simple-golang-app/pkg/http/middleware"
	"simple-golang-app/pkg/logger"

	"github.com/gin-gonic/gin"
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
