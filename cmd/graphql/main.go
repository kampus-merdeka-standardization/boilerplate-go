package main

import (
	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/configs"
	query_pkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/graphql/query"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/graphql/schema"
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
