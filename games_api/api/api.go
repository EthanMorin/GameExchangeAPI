package api

import (
	"games-api/models"
	"games-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct{}

// PostGames implements ServerInterface.
func (*API) PostGames(c *gin.Context) {
	var game models.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.PostGame(&game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// PostUsers implements ServerInterface.
func (*API) PostUsers(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.PostUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// PostExchangesTraderEmailTradeeEmail implements ServerInterface.
func (a *API) PostExchangesTraderEmailTradeeEmail(c *gin.Context, traderEmail string, tradeeEmail string) {
	var exchange models.Exchange
	if err := c.ShouldBindJSON(&exchange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err := services.PostExchange(traderEmail, tradeeEmail, &exchange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetGames implements ServerInterface.
func (*API) GetGames(c *gin.Context) {
	games, err := services.GetGames()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}

// GetGamesId implements ServerInterface.
func (*API) GetGamesId(c *gin.Context, id string) {
	result, err := services.GetGame(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetUsers implements ServerInterface.
func (*API) GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUsersId implements ServerInterface.
func (*API) GetUsersId(c *gin.Context, id string) {
	result, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetExchangeId implements ServerInterface.
func (*API) GetExchangesId(c *gin.Context, id string) {
	result, err := services.GetExchange(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// PatchGamesId implements ServerInterface.
func (*API) PatchGamesId(c *gin.Context, id string) {
	var updatedStatus models.Game
	err := c.ShouldBindJSON(&updatedStatus)
	if err != nil || updatedStatus.Condition == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.PatchGame(id, &updatedStatus)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// PatchUsersId implements ServerInterface.
func (*API) PatchUsersId(c *gin.Context, id string) {
	var updatedUser models.User
	err := c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Binding error": err.Error()})
		return
	}
	services.PatchUser(id, &updatedUser)
	c.JSON(http.StatusOK, &updatedUser)
}

// PatchExchangeId implements ServerInterface.
func (*API) PatchExchangesId(c *gin.Context, id string) {
	var exchange models.PatchExchangesIdJSONBody
	if err := c.ShouldBindJSON(&exchange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if *exchange.Status == "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please enter a valid status"})
		return 
	}
	result, err := services.PatchExchangesId(id, string(*exchange.Status))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteGamesId implements ServerInterface.
func (*API) DeleteGamesId(c *gin.Context, id string) {
	err := services.DeleteGame(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// DeleteUsersId implements ServerInterface.
func (*API) DeleteUsersId(c *gin.Context, id string) {
	err := services.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// DeleteExchangesId implements ServerInterface.
func (a *API) DeleteExchangesId(c *gin.Context, id string) {
	err := services.DeleteExchange(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func NewAPI() *API {
	return &API{}
}
