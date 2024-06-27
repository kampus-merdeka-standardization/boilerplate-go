package main

import (
	"net"

	pinger "simple-golang-app/internal/modules/pinger/grpc"
	pinger_grpc "simple-golang-app/internal/modules/pinger/handler/grpc"
	"simple-golang-app/internal/pkg/configs"
	"simple-golang-app/pkg/logger"

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
