package internal_configs

import (
	"os"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

func LoadConfig() *config {
	appEnv := os.Getenv("PROJECT_ENV")
	if appEnv == "" {
		godotenv.Load(".env")
	} else {
		godotenv.Load(appEnv + ".env")
	}

	consulClientAddress := os.Getenv("CONSUL_CLIENT")

	if consulClientAddress != "" {
		conf, err := loadConfigFromConsul(consulClientAddress)
		if err != nil {
			panic(err)
		}
		return conf
	}

	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
