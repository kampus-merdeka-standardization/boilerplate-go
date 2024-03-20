package configs

type ApiConfig struct {
	Port    string `mapstructure:"PORT"`
	AppEnv  string `mapstructure:"APP_ENV"`
	LogPath string `mapstructure:"LOG_PATH"`
}

type GrpcConfig struct {
	Port    string `mapstructure:"PORT"`
	AppEnv  string `mapstructure:"APP_ENV"`
	LogPath string `mapstructure:"LOG_PATH"`
}

type GraphqlConfig struct {
	Port    string `mapstructure:"PORT"`
	AppEnv  string `mapstructure:"APP_ENV"`
	LogPath string `mapstructure:"LOG_PATH"`
}
