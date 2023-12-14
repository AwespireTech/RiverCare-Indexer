package main

import (
	"github.com/AwespireTech/RiverCare-Indexer/config"
	"github.com/AwespireTech/RiverCare-Indexer/database"
	"github.com/AwespireTech/RiverCare-Indexer/tezos"
)

func main() {
	database.Init(config.DATABASE_URL)
	tezos.Init(config.RPCURL)
	res, err := tezos.GetRiverList(config.FACTORY_BIGMAP)
	if err != nil {
		panic(err)
	}
	for _, river := range res {
		tezos.PrintContractStorage(river)
	}
	tezos.Monitor()

}
