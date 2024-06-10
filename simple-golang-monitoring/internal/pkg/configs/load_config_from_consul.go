package internal_configs

import (
	pkg_consul "simple-golang-monitoring/internal/pkg/consul"
)

func loadConfigFromConsul(address string) (*Config, error) {
	kv, err := pkg_consul.NewConsulKVClient(address)
	if err != nil {
		return nil, err
	}

	otelAddressBytes, err := kv.GetKeyValue("OTEL_ADDRESS", nil)
	if err != nil {
		return nil, err
	}

	appEnvBytes, err := kv.GetKeyValue("APP_ENV", nil)
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

	return &Config{
		OtelAddress: string(otelAddressBytes),
		AppEnv:      string(appEnvBytes),
		AppName:     string(appNameBytes),
		AppPort:     string(appPortBytes),
	}, nil
}
