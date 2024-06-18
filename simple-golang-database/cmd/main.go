package main

import (
	"fmt"
	product_handler "simple-golang-database/internal/modules/product/handler"
	product_postgres "simple-golang-database/internal/modules/product/repository/postgres"
	product_usecase "simple-golang-database/internal/modules/product/usecase"
	internal_configs "simple-golang-database/internal/pkg/configs"
	pkg_db "simple-golang-database/pkg/db"
	pkg_http "simple-golang-database/pkg/http"
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

	productRepository := product_postgres.NewProductRepository(db)
	productUsecase := product_usecase.NewProductUsecase(productRepository)
	product_handler.RegisterProductHandler(srv.Group("/product"), productUsecase)

	if err := srv.Run(fmt.Sprintf(":%s", cfg.AppPort)); err != nil {
		panic(err)
	}
}
