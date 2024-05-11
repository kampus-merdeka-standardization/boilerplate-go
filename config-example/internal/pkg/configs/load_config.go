package internal_configs

import (
	"os"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

func LoadConfig() *config {
	appEnv := os.Getenv("PROJECT_ENV")
	if appEnv == "" {
		godotenv.Load("development.env", ".env")
	} else {
		godotenv.Load(appEnv + ".env")
	}

	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
