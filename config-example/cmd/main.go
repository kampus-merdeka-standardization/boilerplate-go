package main

import (
	internal_configs "config-example/internal/pkg/configs"
	internal_consul "config-example/internal/pkg/consul"
	"fmt"
)

func main() {
	kvClient, err := internal_consul.NewConsulKVClient("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	// config := internal_configs.LoadConfig()
	config := internal_configs.LoadConfigFromConsul(kvClient)

	fmt.Println("App Name : ", config.AppName)
	fmt.Println("App Port : ", config.AppPort)
	fmt.Println("DB Host : ", config.DbHost)
	fmt.Println("DB Username :", config.DbUsername)
	fmt.Println("DB Password : ", config.DbPassword)
}
