package tezos

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintContractStorage(address string) {

	river, err := GetRiverByAddress(address)
	if err != nil {
		log.Panicln(err)
	}
	//Print the river
	buf, err := json.MarshalIndent(river, "", "\t")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("River:")
	fmt.Println(string(buf))

}
