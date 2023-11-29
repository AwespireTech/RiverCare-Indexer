package tezos

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"github.com/AwespireTech/InterfaceForCare-Indexer/config"
)

func GetRiverList(bigmapid string) ([]string, error) {
	query := url.Values{}
	query.Add("active", "true")
	query.Add("select", "key")
	req, err := http.Get(config.TZKT_API_URL + "/bigmaps/" + bigmapid + "/keys" + "?" + query.Encode())
	if err != nil {
		return nil, err
	}
	var keys []string
	err = json.NewDecoder(req.Body).Decode(&keys)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
func GetRiverByAddress(address string) (models.River, error) {
	var river models.River
	client := GetClient()
	addr, err := tezos.ParseAddress(address)
	if err != nil {
		return river, err
	}
	script, err := client.GetContractScript(context.Background(), addr)
	if err != nil {
		return river, err
	}
	val := micheline.NewValue(script.StorageType(), script.Storage)

	river.ID = address
	river.Name, _ = val.GetString("info.name")
	buf, _ := hex.DecodeString(river.Name)
	river.Name = string(buf)
	river.Description, _ = val.GetString("info.description")
	buf, _ = hex.DecodeString(river.Description)
	river.Description = string(buf)

	river.Agreement, _ = val.GetString("agreement_uri")
	buf, _ = hex.DecodeString(river.Agreement)
	river.Agreement = string(buf)

	river.Dataset, _ = val.GetString("dataset_uri")
	buf, _ = hex.DecodeString(river.Dataset)
	river.Dataset = string(buf)

	gen, _ := val.GetInt64("info.generation")
	river.Generation = int(gen)
	river.CreatedTime, _ = val.GetTime("timestamp.create_time")
	river.ExpiredTime, _ = val.GetTime("timestamp.generation_end_time")
	tid, _ := val.GetInt64("stewardship_token.id")
	river.TokenId = int(tid)
	taddr, _ := val.GetAddress("stewardship_token.fa2")
	river.TokenContract = taddr.String()
	river.WalletAddress = address
	//Set Status based on expired time
	if river.ExpiredTime.After(time.Now()) {
		river.Status = models.RIVER_STATUS_ALIVE
	} else {
		river.Status = models.RIVER_STATUS_DEAD
	}
	//Get Events
	event_bigmap, _ := val.GetInt64("event.events")
	events, err := GetAllEventsByBigmap(event_bigmap, river)
	if err == nil {
		river.EventData = events
	}
	//Get Proposals
	proposal_bigmap, _ := val.GetInt64("proposal.proposals")
	proposals, err := GetAllProposalsByBigmap(proposal_bigmap, river)
	if err == nil {
		river.ProposalData = proposals
	}
	owners, err := GetOwners(river.TokenContract, river.TokenId)
	if err != nil {
		return river, err
	}
	river.Owners = owners
	river.OwnersCount = len(owners)

	return river, err
}
