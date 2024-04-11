package rest_client

import "net/http"

func NewRestClient(baseUrl string) *restClient {
	return &restClient{
		baseURL: baseUrl,
		client:  new(http.Client),
	}
}
