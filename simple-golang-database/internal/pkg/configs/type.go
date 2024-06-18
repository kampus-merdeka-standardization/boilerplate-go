package internal_configs

type apiConfig struct {
	AppPort          string `env:"APP_PORT"`
	AppEnv           string `env:"APP_ENV"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresPort     int    `env:"POSTGRES_PORT"`
	PostgresDb       string `env:"POSTGRES_DB"`
}
