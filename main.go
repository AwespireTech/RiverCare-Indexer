package main

import (
	"github.com/AwespireTech/InterfaceForCare-Indexer/config"
	"github.com/AwespireTech/InterfaceForCare-Indexer/database"
	"github.com/AwespireTech/InterfaceForCare-Indexer/tezos"
)

func main() {
	database.Init("mongodb://admin:V0FZoSHM1gUSRcx8UkTMiQ1HYWTAmC11@35.194.208.220:27017")
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
