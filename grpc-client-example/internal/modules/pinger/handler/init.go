package pinger_handler

import (
	"grpc-client-example/internal/modules/pinger/pinger_grpc_gen"

	"google.golang.org/grpc"
)

func NewPingerHandler(conn *grpc.ClientConn) *pingerGrpcHandler {
	c := pinger_grpc_gen.NewPingerServiceClient(conn)

	return &pingerGrpcHandler{
		client: c,
	}
}
