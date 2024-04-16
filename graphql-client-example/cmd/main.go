package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hasura/go-graphql-client"
)

func main() {
	client := graphql.NewClient("http://localhost:8082/graphql", nil)
	// Use client...
	q := query{}
	variables := map[string]any{
		"message": "Palembang",
	}

	ctx := context.Background()
	err := client.Query(ctx, &q, variables)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q.Ping.Message)
}

type query struct {
	Ping struct {
		Message string `graphql:"Message"`
	} `graphql:"ping(message:$message)"`
}
