package tezos

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/AwespireTech/InterfaceForCare-Indexer/config"
)

type Fa2Token struct {
	Owners map[string]int `json:"owners"`
}

func GetOwners(fa2 string, id int) ([]string, error) {
	req, err := http.Get(config.AKASWAP_API_URL + "fa2tokens" + "/" + fa2 + "/" + strconv.Itoa(id))
	log.Println(req.Request.URL)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer req.Body.Close()
	var token Fa2Token
	log.Println(req.ContentLength)
	data := make([]byte, req.ContentLength)
	_, err = req.Body.Read(data)
	if err != nil {
		return nil, err
	}
	log.Println(string(data))
	err = json.Unmarshal(data, &token)
	if err != nil {
		return nil, err
	}
	var owners []string
	for k, v := range token.Owners {
		if v == 0 {
			continue
		}
		owners = append(owners, k)
	}
	return owners, nil

}
