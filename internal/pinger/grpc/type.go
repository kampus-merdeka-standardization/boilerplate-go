package pinger_grpc

import (
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/proto/gen/pinger"
)

type pingerServer struct {
	pinger.UnimplementedPingerServer
}
