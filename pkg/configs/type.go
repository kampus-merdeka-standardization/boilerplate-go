package configs

type ApiConfig struct {
	Port   string `mapstructure:"PORT"`
	AppEnv string `mapstructure:"APP_ENV"`
}

type GrpcConfig struct {
	Port   string `mapstructure:"PORT"`
	AppEnv string `mapstructure:"APP_ENV"`
}

type GraphqlConfig struct {
	Port   string `mapstructure:"PORT"`
	AppEnv string `mapstructure:"APP_ENV"`
}
