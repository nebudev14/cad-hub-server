package routes

import (
	"context"
	"fmt"

	"github.com/NebuDev14/cad-hub-server/db"
	"github.com/NebuDev14/cad-hub-server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var userCollection *mongo.Collection

func StartUsers(group *gin.RouterGroup) {
	userCollection =  db.GetDB().Collection("users")
	group.GET("/:id", GetById)
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid ID: ": err.Error()})
	}
	filter := bson.D{{"_id", objId}}

	var result models.User
	if err := userCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err.Error()})
	}

	c.JSON(http.StatusOK, result)

}
