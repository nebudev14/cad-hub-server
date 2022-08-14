package routes

import (
	"context"
	"github.com/NebuDev14/cad-hub-server/db"
	"github.com/NebuDev14/cad-hub-server/models"
	"github.com/NebuDev14/cad-hub-server/db/helpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var userCollection *mongo.Collection

func GetUserCollection() *mongo.Collection {
	return db.GetDB().Collection("users")
}

func StartUsers(group *gin.RouterGroup) {
	userCollection = db.GetDB().Collection("users")
	group.GET("/:id", GetById)
	group.POST("/:name", CreateUser)
}

func CreateUser(c *gin.Context) {
	name := c.Param("name")
	newUser := models.User{Id: primitive.NewObjectID(), Name: name}

	createNew, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
	}

	c.JSON(http.StatusOK, createNew)
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	result, err := helpers.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest,  gin.H{"error: ": err.Error()})
	}

	c.JSON(http.StatusOK, result)
}
