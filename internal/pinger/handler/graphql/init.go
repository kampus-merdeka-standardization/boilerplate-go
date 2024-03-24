package pinger_graphql

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
)

func NewPingerHandler(ctx *gin.Context) {
	var query struct {
		Query string `json:"query" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&query); err != nil {
		ctx.Error(err)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        newSchema(),
		RequestString: query.Query,
	})

	ctx.JSON(200, httpPkg.Response{
		Message: "Successfully Retrieved Query Response",
		Value:   result,
	})
}

func newSchema() graphql.Schema {
	ql := &pingerGraphql{}

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"ping": &graphql.Field{
				Type:    graphql.String,
				Resolve: ql.Ping,
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		log.Fatal(err)
	}
	return schema
}
