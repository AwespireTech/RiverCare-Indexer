package database

import (
	"context"

	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateStewardHistory(obj models.StewardHistory) error {
	db := GetClient().Database("InterfaceForCare").Collection("stewardshipHistory")
	_, err := db.ReplaceOne(context.Background(), obj, obj, options.Replace().SetUpsert(true))
	return err
}
func UpdateEventHistory(obj models.EventHistory) error {
	db := GetClient().Database("InterfaceForCare").Collection("eventHistory")
	_, err := db.ReplaceOne(context.Background(), obj, obj, options.Replace().SetUpsert(true))
	return err
}
