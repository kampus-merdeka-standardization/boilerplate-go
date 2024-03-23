package main

import (
	"net"

	pinger "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/generated/grpc"
	pinger_grpc "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/grpc"
	product "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/generated/grpc"
	product_grpc "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/handler/grpc"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/configs"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
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
	productServer := product_grpc.NewProductServer()

	pinger.RegisterPingerServiceServer(serverRegistrar, pingerServer)
	product.RegisterProductServiveServer(serverRegistrar, productServer)

	logger.Info("Service is running on Port " + conf.Port)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		logger.Fatal("impossible to server : " + err.Error())
	}
}
