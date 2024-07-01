package rest_client

import "net/http"

type restClient struct {
	baseURL string
	client  *http.Client
}

type RestClient interface {
	SendRequest(method, path string, body any, headers map[string]string) (respBody []byte, code int, err error)
}
