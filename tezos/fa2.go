package tezos

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/AwespireTech/RiverCare-Indexer/config"
)

type AccountInfo struct {
	Address string `json:"address"`
}
type balanceInfo struct {
	AccountInfo AccountInfo `json:"account"`
}

func GetOwners(fa2 string, id int) ([]string, error) {
	query := url.Values{}
	query.Add("token.contract", fa2)
	query.Add("token.tokenId", strconv.Itoa(id))
	query.Add("balance.gt", "0")
	req, err := http.Get(config.TZKT_API_URL + "/tokens/balances" + "?" + query.Encode())
	if err != nil {
		return nil, err
	}
	var balances []balanceInfo
	err = json.NewDecoder(req.Body).Decode(&balances)
	if err != nil {
		return nil, err
	}
	var owners []string
	for _, balance := range balances {
		owners = append(owners, balance.AccountInfo.Address)
	}
	return owners, nil

}
