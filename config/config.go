package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

const (
	// MODE is the mode of the program
	INDEXER = iota
	REBUILD
	DRYRUN
)

var (
	// RPCURL is the url of the tezos node
	RPCURL                string
	AKASWAP_API_URL       string
	TZKT_API_URL          string
	DATABASE_URL          string
	DATABASE_NAME         string
	FACTORY_BIGMAP        string
	TOKEN_METADATA_BIGMAP string
	MODE                  int
)

func init() {
	RPCURL = os.Getenv("RPCURL")
	AKASWAP_API_URL = os.Getenv("AKASWAP_API_URL")
	TZKT_API_URL = os.Getenv("TZKT_API_URL")
	MODE = DRYRUN
	FACTORY_BIGMAP = os.Getenv("FACTORY_BIGMAP")
	TOKEN_METADATA_BIGMAP = os.Getenv("TOKEN_METADATA_BIGMAP")
	var databaseCred string
	if os.Getenv("DATABASE_USERNAME") != "" && os.Getenv("DATABASE_PASSWORD") != "" && os.Getenv("DATABASE_HOST") != "" {
		databaseCred = os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@"
	} else {
		databaseCred = ""
	}
	if os.Getenv("DATABASE_NAME") != "" {
		DATABASE_NAME = os.Getenv("DATABASE_NAME")
	} else {
		DATABASE_NAME = "RiverCare"
	}
	DATABASE_URL = "mongodb://" + databaseCred + os.Getenv("DATABASE_HOST")
}
