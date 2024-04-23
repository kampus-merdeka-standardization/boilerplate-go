package pinger_handler

import (
	"grpc-client-example/internal/modules/pinger/pinger_grpc_gen"
)

type pingerGrpcHandler struct {
	client pinger_grpc_gen.PingerServiceClient
}
