package mock_graphql

import (
	mock_graphql_query "graphql-client/mock/graphql/query"
	mock_graphql_schema "graphql-client/mock/graphql/schema"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func StartGraphqlServer() {
	srv := gin.Default()

	schemaString, err := mock_graphql_schema.String()
	if err != nil {
		panic(err)
	}
	rootQuery := mock_graphql_query.NewRootResolver()
	schema := graphql.MustParseSchema(schemaString, rootQuery)

	srv.POST("/graphql", gin.WrapH(&relay.Handler{Schema: schema}))

	srv.Run(":5000")
}
