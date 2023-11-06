package main

import (
	"time"

	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"github.com/AwespireTech/InterfaceForCare-Indexer/config"
	"github.com/AwespireTech/InterfaceForCare-Indexer/database"
	"github.com/AwespireTech/InterfaceForCare-Indexer/tezos"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Init("mongodb://admin:V0FZoSHM1gUSRcx8UkTMiQ1HYWTAmC11@35.194.208.220:27017")
	config.Init()
	tezos.Init(config.RPCURL)
	tezos.PrintContractStorage("KT19dc6Dmzmo5FAWKiumN76evWV6wcQUmKPq")
	river := models.River{
		ID:            1,
		Name:          "測試的曾文溪",
		Description:   "Description",
		Agreement:     "ipfs://Agreement",
		Dataset:       "ipfs://QmWCqiCjfQtQ8T67K4JheypZxvqoKYBR4SFaHbJd1nQ71R",
		Generation:    0,
		CreatedTime:   time.Unix(0, 0),
		ExpiredTime:   time.Now().AddDate(0, 3, 0),
		Status:        models.RIVER_STATUS_ALIVE,
		TokenId:       -1,
		TokenContract: "KT1QrhY9qyjQ6Q7Q6Q7Q6Q7Q6Q7Q6Q7Q6Q7Q6Q7",
	}
	res, err := database.InsertRiver(river)
	if err != nil {
		panic(err)
	}
	println(res.InsertedID.(int32))

}
