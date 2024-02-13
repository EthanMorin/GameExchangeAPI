package api

import (
	"fmt"
	"games-api/models"
	"games-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type API struct{}

// GetTest implements ServerInterface.
// TODO: Remove after testing im both this and games_api.yaml
func (*API) GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World!")
}

// PostGames implements ServerInterface.
func (*API) PostGames(c *gin.Context) {
	var game models.Game
	
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}
	if game.Condition == nil || game.Name == nil || game.Ownerid == nil || game.Publisher == nil || game.ReleaseYear == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}
	fmt.Print(&game)
	result, err := services.PostGame(&game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert game"})
		return
	}
	c.JSON(http.StatusCreated, result.InsertedID)
}

// PostUsers implements ServerInterface.
func (*API) PostUsers(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}
	if user.Email == nil || user.Password == nil || user.StreetAddress == nil || user.Username == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
	}
	result, err := services.PostUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}
	c.JSON(http.StatusCreated, result.InsertedID)
}

// TODO: add logic to check if traderid, tradeeid, and gameid are valid and if the trader has the game
// PostExchangeTraderidTradeeid implements ServerInterface.
func (*API) PostExchangesTraderidTradeeid(c *gin.Context, traderid string, tradeeid string) {
	// var exchange models.Exchange
	// if err := c.ShouldBindJSON(&exchange); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
	// 	return
	// }
	// if exchange.Gameid == nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
	// 	return
	// }
	// traderId, err := primitive.ObjectIDFromHex(traderid)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid traderid"})
	// 	return
	// }
	// tradeeId, err := primitive.ObjectIDFromHex(tradeeid)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tradeeid"})
	// 	return
	// }
	// status := models.ExchangeStatusPending
	// exchange.Tradeeid = &tradeeId
	// exchange.Traderid = &traderId
	// exchange.Status = &status
	// result, err := services.PostExchange(&exchange)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert exchange"})
	// 	return
	// }
	// c.JSON(http.StatusCreated, result.InsertedID)
}

// GetGames implements ServerInterface.
func (*API) GetGames(c *gin.Context) {
	games, err := services.GetGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get games"})
		return
	}
	c.JSON(http.StatusOK, games)
}

// GetGamesId implements ServerInterface.
func (*API) GetGamesId(c *gin.Context, id string) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	result, err := services.GetGame(objId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetUsers implements ServerInterface.
func (*API) GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUsersId implements ServerInterface.
func (*API) GetUsersId(c *gin.Context, id string) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	result, err := services.GetUser(objId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetExchangeId implements ServerInterface.
func (*API) GetExchangesId(c *gin.Context, id string) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	result, err := services.GetExchange(objId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exchange not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// TODO Test
// PatchGamesId implements ServerInterface.
func (*API) PatchGamesId(c *gin.Context, id string) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var updatedStatus models.Game
	err = c.ShouldBindJSON(&updatedStatus)
	if err != nil || updatedStatus.Condition == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}
	result, err := services.PatchGame(objId, &updatedStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// TODO Test
// PatchUsersId implements ServerInterface.
func (*API) PatchUsersId(c *gin.Context, id string) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var updatedUser models.User
	err = c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}
	services.PatchUser(objId, &updatedUser)
	c.JSON(http.StatusOK, updatedUser)
}

// TODO Impliment
// PatchExchangeId implements ServerInterface.
func (*API) PatchExchangesId(c *gin.Context, id string) {
	var exchange models.Exchange
	// Switch case based off what the user inputed to bind to enum
	err := c.ShouldBindJSON(&exchange)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}
	c.JSON(http.StatusOK, exchange)
}

// DeleteGamesId implements ServerInterface.
func (*API) DeleteGamesId(c *gin.Context, id string) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	services.DeleteGame(objId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
	}
	c.Status(http.StatusNoContent)
}

// DeleteUsersId implements ServerInterface.
func (*API) DeleteUsersId(c *gin.Context, id string) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	services.DeleteUser(objId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.Status(http.StatusNoContent)
}

func NewAPI() *API {
	return &API{}
}
