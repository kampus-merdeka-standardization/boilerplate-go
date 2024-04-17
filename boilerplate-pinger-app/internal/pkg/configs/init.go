package configs

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

func LoadApiConfig() *apiConfig {
	cfg := new(apiConfig)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}

func LoadGrpcConfig() *grpcConfig {
	cfg := new(grpcConfig)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}

func LoadGraphqlConfig() *graphqlConfig {
	cfg := new(graphqlConfig)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
