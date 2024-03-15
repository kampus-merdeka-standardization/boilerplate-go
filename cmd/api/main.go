package main

import (
	pinger_api "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/delivery/api"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
)

func main() {
	srv := httpPkg.NewHTTPServer("debug")
	root := srv.Group("")

	pinger_api.NewPingerController(root)

	srv.Run(":8080")
}
