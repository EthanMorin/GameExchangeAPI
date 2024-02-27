// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Defines values for ExchangeStatus.
const (
	ExchangeStatusAccepted ExchangeStatus = "accepted"
	ExchangeStatusDeclined ExchangeStatus = "declined"
	ExchangeStatusPending  ExchangeStatus = "pending"
)

// Defines values for GameCondition.
const (
	GameConditionFair GameCondition = "fair"
	GameConditionGood GameCondition = "good"
	GameConditionMint GameCondition = "mint"
	GameConditionPoor GameCondition = "poor"
)

// Defines values for PatchExchangesIdJSONBodyStatus.
const (
	PatchExchangesIdJSONBodyStatusAccepted PatchExchangesIdJSONBodyStatus = "accepted"
	PatchExchangesIdJSONBodyStatusDeclined PatchExchangesIdJSONBodyStatus = "declined"
	PatchExchangesIdJSONBodyStatusPending  PatchExchangesIdJSONBodyStatus = "pending"
)

// Defines values for PatchGamesIdJSONBodyStatus.
const (
	PatchGamesIdJSONBodyStatusFair PatchGamesIdJSONBodyStatus = "fair"
	PatchGamesIdJSONBodyStatusGood PatchGamesIdJSONBodyStatus = "good"
	PatchGamesIdJSONBodyStatusMint PatchGamesIdJSONBodyStatus = "mint"
	PatchGamesIdJSONBodyStatusPoor PatchGamesIdJSONBodyStatus = "poor"
)

// Exchange defines model for Exchange.
type Exchange struct {
	Gameid      *primitive.ObjectID `bson:"game_id,omitempty" json:"gameid,omitempty"`
	Id          *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Status      *ExchangeStatus     `json:"status,omitempty"`
	TradeeEmail *string             `json:"tradee_email,omitempty"`
	TraderEmail *string             `json:"trader_email,omitempty"`
}

// ExchangeStatus defines model for Exchange.Status.
type ExchangeStatus string

// Game defines model for Game.
type Game struct {
	Condition   *GameCondition      `json:"condition,omitempty"`
	Id          *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        *string             `json:"name,omitempty"`
	Ownerid     *primitive.ObjectID `bson:"ownerid,omitempty" json:"ownerid,omitempty"`
	Publisher   *string             `json:"publisher,omitempty"`
	ReleaseYear *string             `json:"release_year,omitempty"`
}

// GameCondition defines model for Game.Condition.
type GameCondition string

// User defines model for User.
type User struct {
	Email         *string             `json:"email,omitempty"`
	Id            *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Password      *string             `json:"password,omitempty"`
	StreetAddress *string             `json:"street_address,omitempty"`
	Username      *string             `json:"username,omitempty"`
}

// PatchExchangesIdJSONBody defines parameters for PatchExchangesId.
type PatchExchangesIdJSONBody struct {
	Status *PatchExchangesIdJSONBodyStatus `json:"status,omitempty"`
}

// PatchExchangesIdJSONBodyStatus defines parameters for PatchExchangesId.
type PatchExchangesIdJSONBodyStatus string

// PostExchangesTraderEmailTradeeEmailJSONBody defines parameters for PostExchangesTraderEmailTradeeEmail.
type PostExchangesTraderEmailTradeeEmailJSONBody struct {
	Gameid *string `json:"gameid,omitempty"`
}

// PatchGamesIdJSONBody defines parameters for PatchGamesId.
type PatchGamesIdJSONBody struct {
	Status *PatchGamesIdJSONBodyStatus `json:"status,omitempty"`
}

// PatchGamesIdJSONBodyStatus defines parameters for PatchGamesId.
type PatchGamesIdJSONBodyStatus string

// PatchUsersIdJSONBody defines parameters for PatchUsersId.
type PatchUsersIdJSONBody struct {
	Password *string `json:"password,omitempty"`
}

// PatchExchangesIdJSONRequestBody defines body for PatchExchangesId for application/json ContentType.
type PatchExchangesIdJSONRequestBody PatchExchangesIdJSONBody

// PostExchangesTraderEmailTradeeEmailJSONRequestBody defines body for PostExchangesTraderEmailTradeeEmail for application/json ContentType.
type PostExchangesTraderEmailTradeeEmailJSONRequestBody PostExchangesTraderEmailTradeeEmailJSONBody

// PostGamesJSONRequestBody defines body for PostGames for application/json ContentType.
type PostGamesJSONRequestBody = Game

// PatchGamesIdJSONRequestBody defines body for PatchGamesId for application/json ContentType.
type PatchGamesIdJSONRequestBody PatchGamesIdJSONBody

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = User

// PatchUsersIdJSONRequestBody defines body for PatchUsersId for application/json ContentType.
type PatchUsersIdJSONRequestBody PatchUsersIdJSONBody
