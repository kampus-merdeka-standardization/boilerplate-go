package main

import (
	"log"
	"net"

	pinger_grpc "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/delivery/grpc"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/proto/gen/pinger"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Cannot create listener : %s", err.Error())
	}

	serverRegistrar := grpc.NewServer()
	pingerServer := pinger_grpc.NewPingerServer()

	pinger.RegisterPingerServer(serverRegistrar, pingerServer)

	log.Println("Service is running")
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to server : %s", err.Error())
	}
}
