package main

import (
	"net"

	pinger "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/modules/pinger/grpc"
	pinger_grpc "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/modules/pinger/handler/grpc"
	"github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/pkg/configs"
	"github.com/kampus-merdeka-standardization/boilerplate-pinger-app/pkg/logger"
	"google.golang.org/grpc"
)

func main() {
	conf := configs.LoadGrpcConfig()

	logger.InitLogger(conf.AppEnv, conf.LogPath)

	lis, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		logger.Fatal("Cannot create listener : " + err.Error())
	}

	serverRegistrar := grpc.NewServer()
	pingerServer := pinger_grpc.NewPingerServer()

	pinger.RegisterPingerServiceServer(serverRegistrar, pingerServer)

	logger.Info("Service is running on Port " + conf.Port)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		logger.Fatal("impossible to server : " + err.Error())
	}
}
