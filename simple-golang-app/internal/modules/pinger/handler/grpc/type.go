package pinger_grpc

import (
	pinger "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/modules/pinger/grpc"
)

type pingerServer struct {
	pinger.UnimplementedPingerServiceServer
}
