package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserIdEmail struct {
	Id    *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email *string             `json:"email,omitempty"`
}

type EchangeIds struct {
	Tradeeid *primitive.ObjectID `bson:"tradeeid,omitempty" json:"tradeeid,omitempty"`
	Traderid *primitive.ObjectID `bson:"tradeeid,omitempty" json:"traderid,omitempty"`
}
