package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init(url string) error {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(url).SetServerAPIOptions(serverApi)
	database, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = database.Ping(ctx, nil)
	client = database
	return err
}

func GetClient() *mongo.Client {
	return client
}
func ResetDatabase() error {
	database := client.Database("InterfaceForCare")
	err := database.Drop(context.Background())
	return err
}
func AutoIncreamentId(collection string) int {
	db := client.Database("InterfaceForCare").Collection("autoIncreament")
	var result struct {
		Seq int `bson:"seq"`
	}
	err := db.FindOneAndUpdate(context.Background(), nil, map[string]interface{}{"$inc": map[string]interface{}{"seq": 1}}, options.FindOneAndUpdate().SetUpsert(true)).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result.Seq
}
