package mock_grpc_pinger

import (
	pinger "grpc-client-example/internal/modules/pinger/pinger_grpc_gen"
)

type pingerServer struct {
	pinger.UnimplementedPingerServiceServer
}
