package pinger_grpc

import (
	"context"

	pinger "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/proto/gen/pinger/v1"
)

func (p *pingerServer) Ping(ctx context.Context, req *pinger.PingRequest) (*pinger.PingResponse, error) {
	return &pinger.PingResponse{
		Message: "pong",
	}, nil
}
