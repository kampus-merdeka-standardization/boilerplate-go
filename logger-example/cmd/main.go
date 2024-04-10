package main

import "logger-example/logger"

func main() {
	logger.InitLogger("release", "example-app", "./log/example.log")

	logger.Info("product_file", "create_product", "can't connect to db")
	logger.Debug("product_file", "create_product", "can't connect to db")
	logger.Warn("product_file", "create_product", "can't connect to db")
	logger.Error("product_file", "create_product", "can't connect to db")
	logger.Fatal("product_file", "create_product", "can't connect to db")
}
