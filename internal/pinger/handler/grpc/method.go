package pinger_grpc

import (
	"context"

	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
	pinger "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/proto/gen/pinger/v1"
	"go.uber.org/zap"
)

func (p *pingerServer) Ping(ctx context.Context, req *pinger.PingRequest) (*pinger.PingResponse, error) {
	logger.Debug("Request", zap.Any("req", req))
	return &pinger.PingResponse{
		Message: "pong",
	}, nil
}
