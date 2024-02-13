package data

import (
	"context"
	"games-api/models"
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

func UserCollection() *mongo.Collection {
	return client.Database("Games").Collection("users")
}

func GameCollection() *mongo.Collection {
	return client.Database("Games").Collection("games")
}

func ExchangeCollection() *mongo.Collection {
	return client.Database("Games").Collection("exchanges")
}

// Games
func PostGame(game *models.Game) (*mongo.InsertOneResult, error) {
	result, err := GameCollection().InsertOne(context.Background(), &game)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetGames() (*[]models.Game, error) {
	var games []models.Game
	curser, err := GameCollection().Find(context.Background(), bson.M{})
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
	if err := GameCollection().FindOne(context.Background(), bson.M{"_id": objId}).Decode(&game); err != nil {
		return nil, err
	}
	return &game, nil
}

func PatchGame(objId primitive.ObjectID, game *models.Game) (*mongo.UpdateResult, error) {
	result, err := GameCollection().UpdateOne(context.Background(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"condition": &game.Condition}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteGame(objId primitive.ObjectID) error {
	_, err := GameCollection().DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return err
	}
	return nil
}

// Users
func PostUser(user *models.User) (*mongo.InsertOneResult, error) {
	result, err := UserCollection().InsertOne(context.Background(), &user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUsers() (*[]models.User, error) {
	var users []models.User
	curser, err := UserCollection().Find(context.Background(), bson.M{})
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
	if err := UserCollection().FindOne(context.Background(), bson.M{"_id": objId}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func PatchUser(objId primitive.ObjectID, user *models.User) (*mongo.UpdateResult, error) {
	result, err := UserCollection().UpdateOne(context.Background(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"password": &user.Password}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteUser(objId primitive.ObjectID) error {
	_, err := UserCollection().DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return err
	}
	return nil
}

// Exchanges
// TODO: add more logic to check if users exist and own game
func PostExchange(exchange *models.Exchange) (*mongo.InsertOneResult, error) {
	result, err := ExchangeCollection().InsertOne(context.Background(), &exchange)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetExchange(objId primitive.ObjectID) (*models.Exchange, error) {
	var exchange models.Exchange
	if err := ExchangeCollection().FindOne(context.Background(), bson.M{"_id": objId}).Decode(&exchange); err != nil {
		return nil, err
	}
	return &exchange, nil
}

func PatchExchange(objId primitive.ObjectID, exchangeStatus *models.ExchangeStatus) (*mongo.UpdateResult, error) {
	result, err := ExchangeCollection().UpdateOne(context.Background(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"status": &exchangeStatus}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Helper functions
func GetObjectID(id string) (primitive.ObjectID, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return objId, nil
}
