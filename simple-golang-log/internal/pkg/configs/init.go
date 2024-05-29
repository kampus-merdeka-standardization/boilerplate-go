package configs

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
