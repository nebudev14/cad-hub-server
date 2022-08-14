package models

import (
	"github.com/NebuDev14/cad-hub-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Signature string             `bson:"signature" json:"signature"`
	Projects	[]Project					 `bson:"projects" json:"projects"`
}
