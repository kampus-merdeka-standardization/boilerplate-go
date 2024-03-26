package configs

type apiConfig struct {
	Port             string `env:"API_PORT"`
	AppEnv           string `env:"API_APP_ENV"`
	LogPath          string `env:"API_LOG_PATH"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDb       string `env:"POSTGRES_DB"`
}

type grpcConfig struct {
	Port             string `env:"GRPC_PORT"`
	AppEnv           string `env:"GRPC_APP_ENV"`
	LogPath          string `env:"GRPC_LOG_PATH"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDb       string `env:"POSTGRES_DB"`
}

type graphqlConfig struct {
	Port             string `env:"GRAPHQL_PORT"`
	AppEnv           string `env:"GRAPHQL_APP_ENV"`
	LogPath          string `env:"GRAPHQL_LOG_PATH"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDb       string `env:"POSTGRES_DB"`
}
