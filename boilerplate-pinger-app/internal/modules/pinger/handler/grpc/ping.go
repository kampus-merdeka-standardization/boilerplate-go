package pinger_grpc

import (
	"context"
	"fmt"

	pinger "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/modules/pinger/grpc"
	"github.com/kampus-merdeka-standardization/boilerplate-pinger-app/pkg/logger"
	"go.uber.org/zap"
)

func (p *pingerServer) Ping(ctx context.Context, req *pinger.PingRequest) (*pinger.PingResponse, error) {
	logger.Debug("Request", zap.Any("req", req))
	return &pinger.PingResponse{
		Message: fmt.Sprintf("pong! %s", req.Message),
	}, nil
}
