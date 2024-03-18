package main

import (
	"net"

	pinger_grpc "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/grpc"
	product_grpc "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/handler/grpc"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/configs"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
	pinger "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/proto/gen/pinger/v1"
	product "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/proto/gen/product/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	err := configs.SetEnvVariables("./configs/env/grpc.env")
	if err != nil {
		panic(err)
	}

	var conf configs.GrpcConfig
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	log := logger.NewLogger(conf.AppEnv)

	lis, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		log.Fatal("Cannot create listener : " + err.Error())
	}

	serverRegistrar := grpc.NewServer()
	pingerServer := pinger_grpc.NewPingerServer()
	productServer := product_grpc.NewProductServer()

	pinger.RegisterPingerServiceServer(serverRegistrar, pingerServer)
	product.RegisterProductServiveServer(serverRegistrar, productServer)

	log.Info("Service is running on Port " + conf.Port)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatal("impossible to server : " + err.Error())
	}
}
