package internal_configs

import (
	pkg_consul "config-example/internal/pkg/consul"
	"os"
)

func loadConfigFromConsul(address string) error {
	kv, err := pkg_consul.NewConsulKVClient(address)
	if err != nil {
		return err
	}

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

	os.Setenv("APP_NAME", string(appNameBytes))
	os.Setenv("APP_PORT", string(appPortBytes))
	os.Setenv("DB_USERNAME", string(DbUsernameBytes))
	os.Setenv("DB_PASSWORD", string(DbPasswordBytes))
	os.Setenv("DB_HOST", string(DbHostBytes))

	return nil
}