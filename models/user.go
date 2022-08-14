package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id		primitive.ObjectID		`bson:"_id" json:"id,omitempty"`
	Signature	string					  `bson:"signature" json:"signature"`
}
