package database

import (
	"context"

	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ResetRiverDatabase() error {
	collection := client.Database("InterfaceForCare").Collection("river")
	err := collection.Drop(context.Background())
	return err
}

func InsertRiver(river models.River) (*mongo.InsertOneResult, error) {
	collection := client.Database("InterfaceForCare").Collection("river")
	return collection.InsertOne(context.Background(), river)
}

func UpdateRiver(river models.River) (*mongo.UpdateResult, error) {
	collection := client.Database("InterfaceForCare").Collection("river")
	return collection.UpdateOne(context.Background(), bson.M{"_id": river.ID}, bson.M{"$set": river})

}
