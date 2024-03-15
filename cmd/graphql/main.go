package main

import (
	"log"

	"github.com/gin-gonic/gin"
	pinger_graphql "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/graphql"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
)

func main() {

	srv := httpPkg.NewHTTPServer(gin.DebugMode)

	srv.POST("/pinger", pinger_graphql.NewPingerHandler)

	err := srv.Run(":8082")
	if err != nil {
		log.Fatal(err)
	}
}
