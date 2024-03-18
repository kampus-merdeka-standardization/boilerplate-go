package pinger_grpc

import (
	pinger "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/proto/gen/pinger/v1"
)

type pingerServer struct {
	pinger.UnimplementedPingerServiceServer
}
