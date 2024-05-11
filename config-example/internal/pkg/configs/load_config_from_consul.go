package internal_configs

import pkg_consul "config-example/internal/pkg/consul"

func LoadConfigFromConsul(kv *pkg_consul.KVClient) *config {
	appNameBytes, err := kv.GetKeyValue("APP_NAME", nil)
	if err != nil {
		panic(err)
	}
	appPortBytes, err := kv.GetKeyValue("APP_PORT", nil)
	if err != nil {
		panic(err)
	}
	DbUsernameBytes, err := kv.GetKeyValue("DB_USERNAME", nil)
	if err != nil {
		panic(err)
	}
	DbPasswordBytes, err := kv.GetKeyValue("DB_PASSWORD", nil)
	if err != nil {
		panic(err)
	}
	DbHostBytes, err := kv.GetKeyValue("DB_HOST", nil)
	if err != nil {
		panic(err)
	}

	return &config{
		AppName:    string(appNameBytes),
		AppPort:    string(appPortBytes),
		DbUsername: string(DbUsernameBytes),
		DbPassword: string(DbPasswordBytes),
		DbHost:     string(DbHostBytes),
	}
}
