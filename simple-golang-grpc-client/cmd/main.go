package main

import (
	"context"
	"fmt"
	pinger_client "simple-golang-grpc-client/internal/modules/pinger/client"
	"simple-golang-grpc-client/internal/modules/pinger/pinger_grpc_gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	pHandler := pinger_client.NewPingerHandler(conn)

	res, err := pHandler.Ping(context.Background(), &pinger_grpc_gen.PingRequest{
		Message: "Hello",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
