package configs

type config struct {
	AppEnv  string `env:"APP_ENV"`
	LogPath string `env:"LOG_PATH"`
}
