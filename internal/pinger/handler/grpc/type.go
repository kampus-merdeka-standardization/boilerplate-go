package pinger_grpc

import (
	pinger "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/grpc"
)

type pingerServer struct {
	pinger.UnimplementedPingerServiceServer
}
