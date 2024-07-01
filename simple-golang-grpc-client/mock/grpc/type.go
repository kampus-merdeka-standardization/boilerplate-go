package mock_grpc_pinger

import (
	pinger "simple-golang-grpc-client/internal/modules/pinger/pinger_grpc_gen"
)

type pingerServer struct {
	pinger.UnimplementedPingerServiceServer
}
