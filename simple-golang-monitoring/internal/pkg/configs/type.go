package internal_configs

type Config struct {
	OtelAddress string `env:"OTEL_ADDRESS"`
	AppEnv      string `env:"APP_ENV"`
	AppName     string `env:"APP_NAME"`
	AppPort     string `env:"APP_PORT"`
}

type configLoader struct {
}

type ConfigLoader interface {
	LoadConfig() *Config
}
