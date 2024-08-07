package main

import (
	"fmt"
	product_handler "simple-golang-database/internal/modules/product/handler"
	product_postgres "simple-golang-database/internal/modules/product/repository/postgres"
	product_usecase "simple-golang-database/internal/modules/product/usecase"
	internal_configs "simple-golang-database/internal/pkg/configs"
	pkg_db "simple-golang-database/pkg/db"
	pkg_http "simple-golang-database/pkg/http"
	pkg_http_middleware "simple-golang-database/pkg/http/middleware"
	pkg_logger "simple-golang-database/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := internal_configs.LoadApiConfig()
	db := pkg_db.NewPostgresDB(pkg_db.PostgresDsn{
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		User:     cfg.PostgresUser,
		Password: cfg.PostgresPassword,
		Db:       cfg.PostgresDb,
	})

	srv := pkg_http.NewHTTPServer(cfg.AppEnv)
	pkg_logger.InitLogger(gin.Mode(), "./log/application.log")

	srv.Use(
		gin.Logger(),
		gin.Recovery(),
		pkg_http_middleware.CorsHandler(),
		pkg_http_middleware.ErrorHandler(),
	)

	productRepository := product_postgres.NewProductRepository(db)
	productUsecase := product_usecase.NewProductUsecase(productRepository)
	product_handler.RegisterProductHandler(srv.Group("/product"), productUsecase)

	if err := srv.Run(fmt.Sprintf(":%s", cfg.AppPort)); err != nil {
		panic(err)
	}
}
