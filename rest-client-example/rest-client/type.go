package rest_client

import "net/http"

type restClient struct {
	baseURL string
	client  *http.Client
}
