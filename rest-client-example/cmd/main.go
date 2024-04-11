package main

import (
	"fmt"
	"net/http"
	rest_client "rest-client-example/rest-client"
)

func main() {
	rest := rest_client.NewRestClient("https://api.restful-api.dev")

	getResp, getCode, err := rest.SendRequest(http.MethodGet, "/objects", nil, nil)
	if err != nil {
		panic(err)
	}
	respString := string(getResp)

	fmt.Printf("Get Response : \n")
	fmt.Printf("Code : %d \n", getCode)
	fmt.Printf("Body : %s \n\n", respString)

	body := postBody{
		Name: "Azie Phone",
		Data: bodyData{
			Year:         2024,
			Price:        200.55,
			CpuModel:     "Intel Celeron",
			HardDiskSize: "200 GB",
		},
	}
	postResp, postCode, err := rest.SendRequest(http.MethodPost, "/objects", body, map[string]string{
		"content-type": "application/json",
	})
	if err != nil {
		panic(err)
	}
	respString = string(postResp)

	fmt.Printf("Get Response : \n")
	fmt.Printf("Code : %d \n", postCode)
	fmt.Printf("Body : %s \n\n", respString)
}

type postBody struct {
	Name string   `json:"name"`
	Data bodyData `json:"data"`
}

type bodyData struct {
	Year         int     `json:"year"`
	Price        float64 `json:"price"`
	CpuModel     string  `json:"CPU model"`
	HardDiskSize string  `json:"Hard disk size"`
}
