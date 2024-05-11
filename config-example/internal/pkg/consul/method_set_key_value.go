package internal_consul

import (
	capi "github.com/hashicorp/consul/api"
)

func (client *KVClient) SetKeyValue(key string, value []byte, options *capi.WriteOptions) error {
	_, err := client.kv.Put(&capi.KVPair{
		Key:   key,
		Value: value,
	}, options)

	return err
}
