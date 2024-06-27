package pinger_client

import (
	"simple-golang-grpc-client/internal/modules/pinger/pinger_grpc_gen"
)

type pingerGrpcHandler struct {
	client pinger_grpc_gen.PingerServiceClient
}
