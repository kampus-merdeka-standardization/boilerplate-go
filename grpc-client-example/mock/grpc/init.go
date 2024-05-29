package mock_grpc_pinger

import (
	"grpc-client-example/internal/modules/pinger/pinger_grpc_gen"
	"net"

	"google.golang.org/grpc"
)

func StartGrpcServer() {
	serverRegistrar := grpc.NewServer()
	pServer := pingerServer{}

	pinger_grpc_gen.RegisterPingerServiceServer(serverRegistrar, &pServer)
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	err = serverRegistrar.Serve(lis)
	if err != nil {
		panic(err)
	}
}
