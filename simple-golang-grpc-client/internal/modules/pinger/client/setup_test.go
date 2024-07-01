package pinger_client_test

import (
	mock_grpc_pinger "simple-golang-grpc-client/mock/grpc"
	"sync"
)

func setupGrpcServer() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		mock_grpc_pinger.StartGrpcServer()
	}()
}
