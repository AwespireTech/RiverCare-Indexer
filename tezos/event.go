package tezos

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
	"github.com/AwespireTech/InterfaceForCare-Backend/models"
)

func GetAllEventsByBigmap(bigMapId int64, river models.River) ([]models.Event, error) {
	client := GetClient()
	keys, err := client.ListActiveBigmapKeys(context.Background(), bigMapId)
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
		event.ID = strconv.Itoa(river.ID) + "-" + key.String()
		bigval, err := client.GetActiveBigmapValue(context.Background(), bigMapId, key)
		if err != nil {
			continue
		}
		val := micheline.NewValue(micheline.NewType(info.ValueType), bigval)
		m, _ := val.Map()
		buf, _ := json.MarshalIndent(m, "", "  ")
		fmt.Println(string(buf))
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
		owners, err := GetOwners(river.TokenContract, event.TokenId)

		if err != nil {
			return nil, err
		}
		event.Participants = owners
		event.ParticipantsCount = len(owners)
		gen, _ := val.GetInt64("4")
		event.Generation = int(gen)
		events = append(events, event)
		status, _ := val.GetBool("5")
		if status {
			event.Status = 1
		} else {
			event.Status = 0
		}

	}
	return events, nil
}
