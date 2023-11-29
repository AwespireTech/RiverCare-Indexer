package database

import (
	"context"

	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"github.com/AwespireTech/InterfaceForCare-Indexer/config"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateStewardHistory(obj models.StewardHistory) error {
	db := GetClient().Database(config.DATABASE_NAME).Collection("stewardshipHistory")
	_, err := db.ReplaceOne(context.Background(), obj, obj, options.Replace().SetUpsert(true))
	return err
}
func UpdateEventHistory(obj models.EventHistory) error {
	db := GetClient().Database(config.DATABASE_NAME).Collection("eventHistory")
	_, err := db.ReplaceOne(context.Background(), obj, obj, options.Replace().SetUpsert(true))
	return err
}
