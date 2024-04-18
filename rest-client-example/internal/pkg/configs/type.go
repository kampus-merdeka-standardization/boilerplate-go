package configs

type EnvVariables struct {
	CertFilePath   string `env:"CERTFILE_PATH"`
	KeyFilePath    string `env:"KEYFILE_PATH"`
	CACertFilePath string `env:"CACERTFILE_PATH"`
}
