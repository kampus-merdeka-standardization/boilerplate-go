package configs

type apiConfig struct {
	Port         string `env:"API_PORT"`
	AppEnv       string `env:"API_APP_ENV"`
	LogPath      string `env:"API_LOG_PATH"`
	PostgresConn string `env:"POSTGRES_CONNECTION_STRING"`
}

type grpcConfig struct {
	Port         string `env:"GRPC_PORT"`
	AppEnv       string `env:"GRPC_APP_ENV"`
	LogPath      string `env:"GRPC_LOG_PATH"`
	PostgresConn string `env:"POSTGRES_CONNECTION_STRING"`
}

type graphqlConfig struct {
	Port         string `env:"GRAPHQL_PORT"`
	AppEnv       string `env:"GRAPHQL_APP_ENV"`
	LogPath      string `env:"GRAPHQL_LOG_PATH"`
	PostgresConn string `env:"POSTGRES_CONNECTION_STRING"`
}
