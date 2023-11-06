package config

import (
	"os"
)

const (
	// MODE is the mode of the program
	INDEXER = iota
	REBUILD
	DRYRUN
)

var (
	// RPCURL is the url of the tezos node
	RPCURL          string
	AKASWAP_API_URL string
	MODE            int
)

func Init() {
	RPCURL = os.Getenv("RPCURL")
	AKASWAP_API_URL = os.Getenv("AKASWAP_API_URL")
	MODE = DRYRUN
}
