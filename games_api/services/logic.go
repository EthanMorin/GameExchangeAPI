package services

import (
	"games-api/data"
	"games-api/models"

	"github.com/go-errors/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User
func PostUser(user *models.User) (*models.User, error) {
	if user.Email == nil || user.Password == nil || user.StreetAddress == nil || user.Username == nil {
		return nil, errors.New("All fields are required")
	}
	err := data.PostUser(user)
	if err != nil {
		return nil, errors.New("Couldnt post user")
	}
	return user, nil
}

func GetUsers() (*[]models.User, error) {
	users, err := data.GetUsers()
	if err != nil {
		return nil, errors.New("Couldnt get users")
	}
	return users, nil
}

func GetUser(objId string) (*models.User, error) {
	userObjId, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return nil, errors.New("Invalid object id")
	}
	user, err := data.GetUser(userObjId)
	if err != nil {
		return nil, errors.New("Couldnt get user")
	}
	return user, nil
}

// TODO: TEST THIS
func PatchUser(objId string, user *models.User) (*models.User, error) {
	userObjId, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return nil, errors.New("Invalid object id")
	}
	err = data.PatchUser(userObjId, user)
	if err != nil {
		return nil, errors.New("Couldnt patch user")
	}
	return user, nil
}

func DeleteUser(objId string) error {
	userObjId, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return errors.New("Invalid object id")
	}
	err = data.DeleteUser(userObjId)
	if err != nil {
		return errors.New("Couldnt delete user")
	}
	return nil
}

// Game
func PostGame(game *models.Game) (*models.Game, error) {
	if game.Condition == nil || game.Name == nil || game.Ownerid == nil || game.Publisher == nil || game.ReleaseYear == nil {
		return nil, errors.New("All fields are required")
	}
	err := data.PostGame(game)
	if err != nil {
		return nil, errors.New("Couldnt post game")
	}
	return game, nil
}

func GetGames() (*[]models.Game, error) {
	games, err := data.GetGames()
	if err != nil {
		return nil, errors.New("Couldnt get games")
	}
	return games, nil
}

func GetGame(objId string) (*models.Game, error) {
	gameObjId, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return nil, errors.New("Invalid object id")
	}
	game, err := data.GetGame(gameObjId)
	if err != nil {
		return nil, errors.New("Couldnt get game")
	}
	return game, nil
}

// TODO: TEST THIS
func PatchGame(objId string, game *models.Game) (*models.Game, error) {
	gameObjId, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return nil, errors.New("Invalid object id")
	}
	err = data.PatchGame(gameObjId, game)
	if err != nil {
		return nil, errors.New("Couldnt patch game")
	}
	return game, nil
}

func DeleteGame(objId string) error {
	gameObjId, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return errors.New("Invalid object id")
	}
	err = data.DeleteGame(gameObjId)
	if err != nil {
		return errors.New("Couldnt delete game")
	}
	return nil
}

// Exchange
// TODO: add logic to check if traderid, tradeeid, and gameid are valid and if the trader has the game
func PostExchange(traderId string, tradeeId string, exchange *models.Exchange) (*models.Exchange, error) {
	if exchange.Gameid == nil {
		return nil, errors.New("All fields are required")
	}
	traderObjId, err := primitive.ObjectIDFromHex(traderId)
	if err != nil {
		return nil, errors.New("Invalid traderid")
	}
	tradeeObjId, err := primitive.ObjectIDFromHex(tradeeId)
	if err != nil {
		return nil, errors.New("Invalid tradeeid")
	}
	status := models.ExchangeStatusPending
	exchange.Tradeeid = &tradeeObjId
	exchange.Traderid = &traderObjId
	exchange.Status = &status
	err = data.PostExchange(exchange)
	if err != nil {
		return nil, errors.New("Couldnt post exchange")
	}
	return exchange, nil
}

func GetExchange(objId string) (*models.Exchange, error) {
	exchangeObjId, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return nil, errors.New("Invalid object id")
	}
	exchange, err := data.GetExchange(exchangeObjId)
	if err != nil {
		return nil, errors.New("Couldnt get exchange")
	}
	return exchange, nil
}

// TODO: TEST THIS
func PatchExchangesId(objId string, exchangeStatus *models.ExchangeStatus) (*models.ExchangeStatus, error) {
	exchangeObjId, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return nil, errors.New("Invalid object id")
	}
	err = data.PatchExchange(exchangeObjId, exchangeStatus)
	if err != nil {
		return nil, errors.New("Couldnt patch exchange")
	}
	return exchangeStatus, nil
}