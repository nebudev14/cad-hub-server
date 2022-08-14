package helpers

import (
	"context"
	"fmt"

	"github.com/NebuDev14/cad-hub-server/models"
	"github.com/NebuDev14/cad-hub-server/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserById(id string) models.User {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}
	filter := bson.D{{"_id", objId}}

	var result models.User
	if err := db.GetDB().Collection("users").FindOne(context.TODO(), filter).Decode(&result); err != nil {
		fmt.Println(err)
	}

	return result
}
