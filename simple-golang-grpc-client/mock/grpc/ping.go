package mock_grpc_pinger

import (
	"context"
	"fmt"

	pinger "grpc-client-example/internal/modules/pinger/pinger_grpc_gen"
)

func (p *pingerServer) Ping(ctx context.Context, req *pinger.PingRequest) (*pinger.PingResponse, error) {
	return &pinger.PingResponse{
		Message: fmt.Sprintf("pong! %s", req.Message),
	}, nil
}
