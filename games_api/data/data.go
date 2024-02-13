package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client


func NewDB() (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	// client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://database:27017"))
	if err != nil {
		return err
	}
	return nil
}

func UserCollection() *mongo.Collection {
	return client.Database("Games").Collection("users")
}

func GameCollection() *mongo.Collection {
	return client.Database("Games").Collection("games")
}

func ExchangeCollection() *mongo.Collection {
	return client.Database("Games").Collection("exchanges")
}

