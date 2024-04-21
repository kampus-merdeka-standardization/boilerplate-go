package main

import (
	"fmt"
	pinger_client "graphql-client/internal/modules/pinger/client"

	"github.com/hasura/go-graphql-client"
)

func main() {
	client := graphql.NewClient("http://localhost:8082/graphql", nil)

	pingerClient := pinger_client.NewPingerClient(client)

	msg, err := pingerClient.GetMessage("Palembang")
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
}
