package internal_configs

type Config struct {
	// APP_NAME
	// APP_PORT
	// DB_USERNAME
	// DB_PASSWORD
	// DB_HOST
	AppName    string `env:"APP_NAME"`
	AppPort    string `env:"APP_PORT"`
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
	DbHost     string `env:"DB_HOST"`
}

type configLoader struct {
}

type ConfigLoader interface {
	LoadConfig() *Config
}
