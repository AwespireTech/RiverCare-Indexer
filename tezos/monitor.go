package tezos

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"blockwatch.cc/tzgo/rpc"
	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"github.com/AwespireTech/InterfaceForCare-Indexer/config"
	"github.com/AwespireTech/InterfaceForCare-Indexer/database"
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
func FullUpdate() error {
	list, err := GetRiverList(config.FACTORY_BIGMAP)
	if err != nil {
		return err
	}
	for _, address := range list {
		river, err := GetRiverByAddress(address)
		if err != nil {
			return err
		}
		err = UpdateRiver(river)
		if err != nil {
			return err
		}
	}
	return nil
}
func Monitor() {
	client := GetClient()
	mon := rpc.NewBlockHeaderMonitor()
	defer mon.Close()
	client.MonitorBlockHeader(context.Background(), mon)
	for {
		block, err := mon.Recv(context.Background())
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(block.Hash)

		// err = FullUpdate()
		// if err != nil {
		// 	log.Panicln(err)
		// }
	}

}
func UpdateRiver(river models.River) error {
	//Parse EventData to Id
	var eventIds []string
	for _, event := range river.EventData {
		eventIds = append(eventIds, event.ID)
	}
	river.Events = eventIds
	//Parse ProposalData to Id
	var proposalIds []string
	for _, proposal := range river.ProposalData {
		proposalIds = append(proposalIds, proposal.ID)
	}
	river.Proposals = proposalIds
	//Update ProposalData
	for _, proposal := range river.ProposalData {
		_, err := database.UpdateProposal(proposal)
		if err != nil {
			return err
		}
	}
	//Update EventData
	for _, event := range river.EventData {
		_, err := database.UpdateEvent(event)
		if err != nil {
			return err
		}
	}
	//Update TokenHistory
	for _, owner := range river.Owners {
		hist := models.StewardHistory{
			RiverId:       river.ID,
			User:          owner,
			TokenContract: river.TokenContract,
			TokenId:       river.TokenId,
			Generation:    river.Generation,
		}
		err := database.UpdateStewardHistory(hist)
		if err != nil {
			return err
		}
	}
	return nil

}
