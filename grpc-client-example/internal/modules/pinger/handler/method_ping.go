package pinger_handler

import (
	"context"
	"grpc-client-example/internal/modules/pinger/pinger_grpc_gen"
)

func (p *pingerGrpcHandler) Ping(ctx context.Context, pingerRes *pinger_grpc_gen.PingRequest) (*pinger_grpc_gen.PingResponse, error) {
	res, err := p.client.Ping(ctx, pingerRes)
	if err != nil {
		return nil, err
	}

	return res, nil
}
