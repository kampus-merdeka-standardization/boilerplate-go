package internal_configs

type config struct {
	AppName string `env:"APP_NAME"`
	AppEnv  string `env:"APP_ENV"`
	AppPort string `env:"APP_PORT"`
	LogPath string `env:"LOG_PATH"`
}
