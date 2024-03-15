package pinger_grpc

import (
	"context"

	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/proto/gen/pinger"
)

func (p *pingerServer) Ping(ctx context.Context, req *pinger.PingRequest) (*pinger.PingResponse, error) {
	return &pinger.PingResponse{
		Response: "Pong From Ping's " + req.Request,
	}, nil
}
