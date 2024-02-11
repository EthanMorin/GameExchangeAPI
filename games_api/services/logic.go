package services

// TODO move logic from api/api.go to here

import (
	"games-api/data"
	"games-api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Games
func PostGame(game *models.Game) (*mongo.InsertOneResult, error) {
	result, err := data.GameCollection().InsertOne(context.Background(), &game)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetGames() (*[]models.Game, error) {
	var games []models.Game
	curser, err := data.GameCollection().Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err := curser.All(context.Background(), &games); err != nil {
		return nil, err
	}
	return &games, nil
}

func GetGame(objId primitive.ObjectID) (*models.Game, error) {
	var game models.Game
	if err := data.GameCollection().FindOne(context.Background(), bson.M{"_id": objId}).Decode(&game); err != nil {
		return nil, err
	}
	return &game, nil
}

func PatchGame(objId primitive.ObjectID, game *models.Game) (*mongo.UpdateResult, error) {
	result, err := data.GameCollection().UpdateOne(context.Background(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"condition": &game.Condition}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteGame(objId primitive.ObjectID) error {
	_, err := data.GameCollection().DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return err
	}
	return nil
}

// Users
func PostUser(user *models.User) (*mongo.InsertOneResult, error) {
	result, err := data.UserCollection().InsertOne(context.Background(), &user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUsers() (*[]models.User, error) {
	var users []models.User
	curser, err := data.UserCollection().Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err := curser.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return &users, nil
}

func GetUser(objId primitive.ObjectID) (*models.User, error) {
	var user models.User
	if err := data.UserCollection().FindOne(context.Background(), bson.M{"_id": objId}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func PatchUser(objId primitive.ObjectID, user *models.User) (*mongo.UpdateResult, error) {
	result, err := data.UserCollection().UpdateOne(context.Background(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"password": &user.Password}})
	if err != nil {
		return nil, err
	}
	return result, nil
}


func DeleteUser(objId primitive.ObjectID) error {
	_, err := data.UserCollection().DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return err
	}
	return nil
}

// Exchanges
// TODO: add more logic to check if users exist and own game
func PostExchange(exchange *models.Exchange) (*mongo.InsertOneResult, error) {
	result, err := data.ExchangeCollection().InsertOne(context.Background(), &exchange)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetExchange(objId primitive.ObjectID) (*models.Exchange, error) {
	var exchange models.Exchange
	if err := data.ExchangeCollection().FindOne(context.Background(), bson.M{"_id": objId}).Decode(&exchange); err != nil {
		return nil, err
	}
	return &exchange, nil
}

// Helper functions
func GetObjectID(id string) (primitive.ObjectID, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return objId, nil
}