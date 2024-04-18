package main

import (
	"context"
	"fmt"
	pinger_handler "grpc-client-example/internal/modules/pinger/handler"
	"grpc-client-example/internal/modules/pinger/pinger_grpc_gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	pHandler := pinger_handler.NewPingerHandler(conn)

	res, err := pHandler.Ping(context.Background(), &pinger_grpc_gen.PingRequest{
		Message: "Hello",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
