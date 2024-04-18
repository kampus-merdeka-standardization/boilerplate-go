package main

import (
	"fmt"
	"net/http"
	rest_client "rest-client-example/internal/modules/rest-client"
	"rest-client-example/internal/pkg/configs"
)

func main() {
	cfg := configs.LoadEnv()

	client, err := rest_client.NewRestClient("https://api.restful-api.dev")
	if err != nil {
		panic(err)
	}

	tlsClient, err := rest_client.NewRestClientWithTls("https://api.restful-api.dev", cfg.CertFilePath, cfg.KeyFilePath, cfg.CACertFilePath)
	if err != nil {
		panic(err)
	}

	res, code, err := client.SendRequest(http.MethodGet, "/objects/2", nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Code : ", code)
	fmt.Println(string(res))

	res, code, err = tlsClient.SendRequest(http.MethodGet, "/objects/5", nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Code : ", code)
	fmt.Println(string(res))

	type bodyData struct {
		Year         int     `json:"year"`
		Price        float64 `json:"price"`
		CPUModel     string  `json:"CPU model"`
		HardDiskSize string  `json:"Hard disk size"`
	}
	type bodyRequest struct {
		Name string   `json:"name"`
		Data bodyData `json:"data"`
	}

	reqBody := bodyRequest{
		Name: "Lenovo Thinkpad 5",
		Data: bodyData{
			Year:         2021,
			Price:        599.99,
			CPUModel:     "Intel Core i5 Gen 10",
			HardDiskSize: "1 TB",
		},
	}

	res, code, err = client.SendRequest(http.MethodPost, "/objects", reqBody, map[string]string{
		"content-type": "application/json",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Code : ", code)
	fmt.Println(string(res))

	res, code, err = tlsClient.SendRequest(http.MethodPost, "/objects", reqBody, map[string]string{
		"content-type": "application/json",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Code : ", code)
	fmt.Println(string(res))
}
