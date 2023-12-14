package tezos

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
	"github.com/AwespireTech/RiverCare-Backend/models"
	"github.com/AwespireTech/RiverCare-Indexer/config"
)

func getEventTokenMetadata(tokenId int) (map[string]string, error) {
	bigMapId := config.TOKEN_METADATA_BIGMAP
	req, err := http.Get(config.TZKT_API_URL + "/bigmaps/" + bigMapId + "/keys/" + strconv.Itoa(tokenId))
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	data = data["value"].(map[string]interface{})
	data = data["token_info"].(map[string]interface{})

	res := make(map[string]string)
	buf, err := hex.DecodeString(data["name"].(string))
	if err != nil {
		return nil, err
	}
	res["name"] = string(buf)
	buf, err = hex.DecodeString(data["description"].(string))
	if err != nil {
		return nil, err
	}
	res["description"] = string(buf)
	buf, err = hex.DecodeString(data["displayUri"].(string))
	if err != nil {
		return nil, err
	}
	res["displayUri"] = string(buf)

	return res, nil
}

func GetAllEventsByBigmap(bigMapId int64, river models.River) ([]models.Event, error) {
	client := GetClient()

	query := url.Values{}
	query.Add("active", "true")
	query.Add("select", "key,hash")
	query.Add("limit", "10000")
	req, err := http.Get(config.TZKT_API_URL + "/bigmaps/" + strconv.FormatInt(bigMapId, 10) + "/keys" + "?" + query.Encode())
	if err != nil {
		return nil, err
	}

	var keys []map[string]string
	err = json.NewDecoder(req.Body).Decode(&keys)
	if err != nil {
		return nil, err
	}

	info, err := client.GetActiveBigmapInfo(context.Background(), bigMapId)
	if err != nil {
		return nil, err
	}

	var events []models.Event
	for _, key := range keys {
		var event models.Event

		hash, err := tezos.ParseExprHash(key["hash"])
		if err != nil {
			continue
		}
		event.ID = river.ID + "-" + key["key"]

		bigval, err := client.GetActiveBigmapValue(context.Background(), bigMapId, hash)
		if err != nil {
			continue
		}
		val := micheline.NewValue(micheline.NewType(info.ValueType), bigval)
		// m, _ := val.Map()
		// buf, _ := json.MarshalIndent(m, "", "  ")
		// fmt.Println(string(buf))
		amount, isLimited := val.GetInt64("0")
		if isLimited {
			event.Amount = int(amount)
		} else {
			event.Amount = -1
		}
		event.CreatedTime, _ = val.GetTime("3")
		host, _ := val.GetAddress("6")
		event.Host = host.String()
		appr, _ := val.GetValue("1")
		approvers, _ := appr.([]interface{})
		approvals := make([]string, 0)
		for _, approver := range approvers {
			tmp := approver.(tezos.Address)
			approvals = append(approvals, tmp.String())
		}
		event.Approvals = approvals
		event.ApprovalsCount = len(event.Approvals)
		tid, _ := val.GetBig("7")
		event.TokenId = int(tid.Int64())
		event.TokenContract = river.TokenContract

		data, err := getEventTokenMetadata(event.TokenId)
		if err != nil {
			return nil, err
		}
		event.Name = data["name"]
		event.Description = data["description"]
		event.ImageUri = data["displayUri"]

		owners, err := GetOwners(river.TokenContract, event.TokenId)
		if err != nil {
			return nil, err
		}
		event.Participants = owners
		event.ParticipantsCount = len(owners)
		gen, _ := val.GetInt64("4")
		event.Generation = int(gen)
		status, _ := val.GetBool("5")
		if status {
			event.Status = 1
		} else {
			event.Status = 0
		}
		event.Editions = int(amount) + len(owners)
		events = append(events, event)

	}
	return events, nil
}
