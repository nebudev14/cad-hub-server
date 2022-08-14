package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	Id    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Users []User             `bson:"users" json:"users"`
}
