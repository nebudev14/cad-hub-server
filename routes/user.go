package routes

import (
	"context"
	"fmt"

	"github.com/NebuDev14/cad-hub-server/db"
	"github.com/NebuDev14/cad-hub-server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func StartUsers(group *gin.RouterGroup) {
	userCollection =  db.GetDB().Collection("users")
	group.GET("/:id", GetById)
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	filter := bson.D{{"_id", id}}

	var result models.User
	if err := userCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		fmt.Println(err)
	}

}
