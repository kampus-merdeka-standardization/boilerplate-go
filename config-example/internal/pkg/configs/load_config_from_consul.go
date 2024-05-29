package internal_configs

import (
	pkg_consul "config-example/internal/pkg/consul"
)

func loadConfigFromConsul(address string) (*config, error) {
	kv, err := pkg_consul.NewConsulKVClient(address)
	if err != nil {
		return nil, err
	}

	appNameBytes, err := kv.GetKeyValue("APP_NAME", nil)
	if err != nil {
		return nil, err
	}
	appPortBytes, err := kv.GetKeyValue("APP_PORT", nil)
	if err != nil {
		return nil, err
	}
	DbUsernameBytes, err := kv.GetKeyValue("DB_USERNAME", nil)
	if err != nil {
		return nil, err
	}
	DbPasswordBytes, err := kv.GetKeyValue("DB_PASSWORD", nil)
	if err != nil {
		return nil, err
	}
	DbHostBytes, err := kv.GetKeyValue("DB_HOST", nil)
	if err != nil {
		return nil, err
	}

	return &config{
		AppName:    string(appNameBytes),
		AppPort:    string(appPortBytes),
		DbUsername: string(DbUsernameBytes),
		DbPassword: string(DbPasswordBytes),
		DbHost:     string(DbHostBytes),
	}, nil
}
