package tezos

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"blockwatch.cc/tzgo/micheline"
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
		event.Host, _ = val.GetString("6")
		tid, _ := val.GetBig("7")
		event.TokenId = int(tid.Int64())
		events = append(events, event)
	}
	return events, nil
}