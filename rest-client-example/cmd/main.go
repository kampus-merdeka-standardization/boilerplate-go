package main

import (
	"fmt"
	"net/http"
	rest_client "rest-client-example/internal/modules/rest-client"
	"rest-client-example/internal/pkg/configs"
)

func main() {
	cfg := configs.LoadEnv()

	client := rest_client.NewRestClient("https://api.restful-api.dev")

	res, code, err := client.SendRequest(http.MethodGet, "/objects/2", nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Code : ", code)
	fmt.Println(string(res))

	tlsClient := rest_client.NewRestClientWithTls("https://api.restful-api.dev", cfg.CertFilePath, cfg.KeyFilePath, cfg.CACertFilePath)

	res, code, err = tlsClient.SendRequest(http.MethodGet, "/objects/5", nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Code : ", code)
	fmt.Println(string(res))
}
