package database

import (
	"context"

	"github.com/AwespireTech/RiverCare-Backend/models"
	"github.com/AwespireTech/RiverCare-Indexer/config"
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
