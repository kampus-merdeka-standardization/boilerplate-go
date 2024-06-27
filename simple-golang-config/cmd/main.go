package main

import (
	"fmt"
	internal_configs "simple-golang-config/internal/pkg/configs"
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
