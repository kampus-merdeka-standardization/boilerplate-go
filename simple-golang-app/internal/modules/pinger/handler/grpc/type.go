package pinger_grpc

import (
	pinger "simple-golang-app/internal/modules/pinger/grpc"
)

type pingerServer struct {
	pinger.UnimplementedPingerServiceServer
}
