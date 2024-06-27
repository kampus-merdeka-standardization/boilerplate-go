package main

import (
	"simple-golang-app/internal/pkg/configs"
	query_pkg "simple-golang-app/internal/pkg/graphql/query"
	"simple-golang-app/internal/pkg/graphql/schema"
	httpPkg "simple-golang-app/pkg/http"
	"simple-golang-app/pkg/http/middleware"
	"simple-golang-app/pkg/logger"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	conf := configs.LoadGraphqlConfig()

	srv := httpPkg.NewHTTPServer(conf.AppEnv)

	logger.InitLogger(conf.AppEnv, conf.LogPath)

	srv.Use(middleware.LogHandler(), gin.Recovery())
	srv.Use(middleware.CorsHandler())

	schemaString, err := schema.String()
	if err != nil {
		logger.Fatal(err.Error())
	}
	rootQuery := query_pkg.NewRootResolver()
	schema := graphql.MustParseSchema(schemaString, rootQuery)

	srv.POST("/graphql", gin.WrapH(&relay.Handler{Schema: schema}))

	logger.Info("Running on Port " + conf.Port)
	if err := srv.Run(":" + conf.Port); err != nil {
		logger.Fatal(err.Error())
	}
}
