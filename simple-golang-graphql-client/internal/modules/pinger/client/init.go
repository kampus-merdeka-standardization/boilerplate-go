package pinger_client

import "github.com/hasura/go-graphql-client"

func NewPingerClient(client *graphql.Client) *pingerClient {
	return &pingerClient{client: client}
}
