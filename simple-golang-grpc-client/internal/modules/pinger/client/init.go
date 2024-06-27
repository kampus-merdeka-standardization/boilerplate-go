package pinger_client

import (
	"simple-golang-grpc-client/internal/modules/pinger/pinger_grpc_gen"

	"google.golang.org/grpc"
)

func NewPingerHandler(conn *grpc.ClientConn) *pingerGrpcHandler {
	c := pinger_grpc_gen.NewPingerServiceClient(conn)

	return &pingerGrpcHandler{
		client: c,
	}
}
