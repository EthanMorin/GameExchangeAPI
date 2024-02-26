package data

import (
	"context"
	"games-email/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func NewDB() error {
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

func userCollection() *mongo.Collection {
	return client.Database("Games").Collection("users")
}

func GetUserEmail(objId primitive.ObjectID) (string, error) {
	var user models.UserIdEmail
	if err := userCollection().FindOne(context.Background(), bson.M{"_id": objId}).Decode(&user); err != nil {
		return "", err
	}
	return *user.Email, nil
}