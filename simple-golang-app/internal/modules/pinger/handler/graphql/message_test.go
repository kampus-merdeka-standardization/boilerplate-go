package pinger_resolver_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	queryPkg "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/pkg/graphql/query"
	"github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/pkg/graphql/schema"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/pkg/http"
	"github.com/kampus-merdeka-standardization/boilerplate-pinger-app/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func setupTest() *gin.Engine {
	srv := httpPkg.NewHTTPServer("test")

	schemaString, err := schema.String()
	if err != nil {
		logger.Fatal(err.Error())
	}
	rootQuery := queryPkg.NewRootResolver()
	schema := graphql.MustParseSchema(schemaString, rootQuery)

	srv.POST("/graphql", gin.WrapH(&relay.Handler{Schema: schema}))

	return srv
}
func TestMessageGraphQL(t *testing.T) {
	router := setupTest()

	bodyReader := strings.NewReader(
		`{"query":"query Ping {\n    ping(message:\"Azie\"){\n        Message\n    }\n}\n","variables":{}}`)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/graphql", bodyReader)

	router.ServeHTTP(res, req)

	var resBody httpPkg.Response
	json.Unmarshal(res.Body.Bytes(), &resBody)
	assert.Equal(t, http.StatusOK, res.Code)
	log.Println(res.Body.String())
}
