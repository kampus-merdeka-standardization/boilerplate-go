package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
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

	srv.Handle(http.MethodGet, "", gin.WrapH(handler.New(&handler.Config{
		Schema:     graphqlSchema(),
		Pretty:     true,
		Playground: true,
	})))

	logger.Info("Running on Port " + conf.Port)
	if err := srv.Run(":" + conf.Port); err != nil {
		logger.Fatal(err.Error())
	}
}

func graphqlSchema() *graphql.Schema {
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name: "Root",
		Fields: graphql.Fields{
			"ping": pinger_graphql.NewField(),
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: obj,
	})
	if err != nil {
		logger.Fatal(err.Error())
	}

	return &schema
}
