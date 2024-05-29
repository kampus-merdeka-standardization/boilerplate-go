package pinger_client

import "github.com/hasura/go-graphql-client"

type pingerClient struct {
	client *graphql.Client
}

type query struct {
	Ping struct {
		Message string `graphql:"Message"`
	} `graphql:"ping(message:$message)"`
}
