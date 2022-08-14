package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Signature string             `bson:"signature" json:"signature"`
	Projects	[]Project					 `bson:"projects" json:"projects"`
}
