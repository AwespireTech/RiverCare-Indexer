package database

import (
	"context"

	"github.com/AwespireTech/RiverCare-Backend/models"
	"github.com/AwespireTech/RiverCare-Indexer/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ResetRiverDatabase() error {
	collection := client.Database(config.DATABASE_NAME).Collection("river")
	err := collection.Drop(context.Background())
	return err
}

func InsertRiver(river models.River) (*mongo.InsertOneResult, error) {
	collection := client.Database(config.DATABASE_NAME).Collection("river")
	return collection.InsertOne(context.Background(), river)
}

func UpdateRiver(river models.River) (*mongo.UpdateResult, error) {
	collection := client.Database(config.DATABASE_NAME).Collection("river")
	return collection.ReplaceOne(context.Background(), bson.M{"_id": river.ID}, river, options.Replace().SetUpsert(true))
}
func UpdateProposal(proposal models.Proposal) (*mongo.UpdateResult, error) {
	collection := client.Database(config.DATABASE_NAME).Collection("proposal")
	return collection.ReplaceOne(context.Background(), bson.M{"_id": proposal.ID}, proposal, options.Replace().SetUpsert(true))
}
func UpdateEvent(event models.Event) (*mongo.UpdateResult, error) {
	collection := client.Database(config.DATABASE_NAME).Collection("event")
	return collection.ReplaceOne(context.Background(), bson.M{"_id": event.ID}, event, options.Replace().SetUpsert(true))
}
