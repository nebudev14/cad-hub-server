package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Change struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Project   Project            `bson:"project" json:"project"`
	User      User               `bson:"user" json:"user"`
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt,omitempty"`
}
