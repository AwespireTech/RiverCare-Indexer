package tezos

import (
	"context"
	"encoding/hex"
	"errors"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
	"github.com/AwespireTech/RiverCare-Backend/models"
)

func decodeContent(content interface{}) (int, map[string]interface{}, error) {
	code := 0
	var layer map[string]interface{}
	data := content
	for i := 0; i < 3; i++ {
		layer = data.(map[string]interface{})
		if layer["@or_0"] != nil {
			code = code<<1 + 0
			data = layer["@or_0"]
		} else {
			code = code<<1 + 1
			data = layer["@or_1"]
		}
	}
	m := map[string]interface{}{
		"dataset":   "",
		"agreement": "",
		"amount":    int64(0),
		"to":        "",
	}
	t := -1
	switch code {
	case 4:
		t = models.PROPOSAL_TYPE_DATASET
		buf, _ := hex.DecodeString(data.(string))
		m["dataset"] = string(buf)
	case 3:
		t = models.PROPOSAL_TYPE_AGREEMENT
		buf, _ := hex.DecodeString(data.(string))
		m["agreement"] = string(buf)
	case 2:
		t = models.PROPOSAL_TYPE_TRANSFER
		d := data.(map[string]interface{})
		m["amount"] = d["0"].(tezos.Z).Int64()
		m["to"] = d["1"].(tezos.Address).String()
	default:
		return -1, nil, errors.New("unknown type")
	}

	return t, m, nil
}

func GetAllProposalsByBigmap(bigMapId int64, river models.River) ([]models.Proposal, error) {
	client := GetClient()
	keys, err := client.ListActiveBigmapKeys(context.Background(), bigMapId)
	if err != nil {
		return nil, err
	}

	info, err := client.GetActiveBigmapInfo(context.Background(), bigMapId)
	if err != nil {
		return nil, err
	}

	var proposals []models.Proposal
	for _, key := range keys {
		var proposal models.Proposal

		proposal.ID = river.ID + "-" + key.String()
		bigval, err := client.GetActiveBigmapValue(context.Background(), bigMapId, key)
		if err != nil {
			continue
		}
		val := micheline.NewValue(micheline.NewType(info.ValueType), bigval)
		// m, _ := val.Map()
		// buf, _ := json.MarshalIndent(m, "", "  ")
		// fmt.Println(string(buf))

		content, _ := val.GetValue("1")
		t, data, err := decodeContent(content)
		if err != nil {
			continue
		}

		proposal.TransactionType = t
		proposal.Agreement = data["agreement"].(string)
		proposal.Dataset = data["dataset"].(string)
		proposal.TargetAddress = data["to"].(string)
		proposal.TransferMutez = data["amount"].(int64)

		proposer, _ := val.GetAddress("5")
		proposal.ProposerAddress = proposer.String()
		status, _ := val.GetBool("4")
		if status {
			proposal.Status = models.PROPOSAL_STATUS_APPROVED
		} else {
			proposal.Status = models.PROPOSAL_STATUS_PROPOSED
		}
		proposal.CreatedTime, _ = val.GetTime("2")
		proposal.ExpiredTime = river.ExpiredTime
		appr, _ := val.GetValue("0")
		approvers, _ := appr.([]interface{})
		approvals := make([]string, 0)
		for _, approver := range approvers {
			tmp := approver.(tezos.Address)
			approvals = append(approvals, tmp.String())
		}
		proposal.Approvals = approvals
		proposal.ApprovalsCount = len(proposal.Approvals)
		gen, _ := val.GetInt64("3")
		proposal.Generation = int(gen)
		proposals = append(proposals, proposal)
	}
	return proposals, nil
}
