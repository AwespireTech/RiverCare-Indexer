package tezos

import (
	"blockwatch.cc/tzgo/rpc"
)

var client *rpc.Client

func Init(url string) error {
	_client, err := rpc.NewClient(url, nil)
	if err != nil {
		return err
	}
	client = _client
	return nil
}
func GetClient() *rpc.Client {
	return client
}
