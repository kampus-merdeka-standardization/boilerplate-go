package main

import (
	internal_configs "config-example/internal/pkg/configs"
	"fmt"
)

func main() {
	configLoader := internal_configs.NewConfigLoader()
	config := configLoader.LoadConfig()

	fmt.Println("App Name : ", config.AppName)
	fmt.Println("App Port : ", config.AppPort)
	fmt.Println("DB Host : ", config.DbHost)
	fmt.Println("DB Username :", config.DbUsername)
	fmt.Println("DB Password : ", config.DbPassword)
}
