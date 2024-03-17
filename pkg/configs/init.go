package configs

import "github.com/spf13/viper"

func SetEnvVariables(envPath string) error {
	// Load environment variables from .env file
	viper.SetConfigFile(envPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Set environment variables
	viper.AutomaticEnv()
	return nil
}
