package main

import (
	"github.com/gin-gonic/gin"
	pinger_api "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/api"
	product_api "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/handler/api"
	product_postgres "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/repository/postgres"
	product_usecase "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/usecase"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/configs"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/db"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http/middleware"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
)

func main() {
	conf := configs.LoadApiConfig()

	srv := httpPkg.NewHTTPServer(conf.AppEnv)

	logger.InitLogger(conf.AppEnv, conf.LogPath)

	srv.Use(middleware.LogHandler(), gin.Recovery())
	srv.Use(middleware.CorsHandler())

	pgDb := db.NewPostgresDB(conf.PostgresHost, conf.PostgresUser, conf.PostgresPassword, conf.Port, conf.PostgresDb)

	pinger_api.NewPingerController(srv.Group("/ping"))

	productRepository := product_postgres.NewProductRepository(pgDb)
	productUsecase := product_usecase.NewProductUsecase(productRepository)
	product_api.NewProductController(srv.Group("/product"), productUsecase)

	logger.Info("Running on Port " + conf.Port)
	if err := srv.Run(":" + conf.Port); err != nil {
		logger.Fatal(err.Error())
	}
}
