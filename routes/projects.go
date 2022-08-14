package routes

import (
	"context"
	"fmt"

	"net/http"

	"github.com/NebuDev14/cad-hub-server/db"
	"github.com/NebuDev14/cad-hub-server/db/helpers"
	"github.com/NebuDev14/cad-hub-server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var projectCollection *mongo.Collection

func GetProjectCollection() *mongo.Collection {
	return db.GetDB().Collection("projects")
}

type CreateProjectReq struct {
	UserId		string
	Name			string
}

func StartProjects(group *gin.RouterGroup) {
	projectCollection = db.GetDB().Collection("projects")
	group.POST("", CreateProject)
}

func CreateProject(c *gin.Context) {
	var reqBody CreateProjectReq
	if err := c.Copy().ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
	}

	user, err := helpers.GetUserById(reqBody.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest,  gin.H{"error: ": err.Error()})
	}

	newProject := models.Project{Id: primitive.NewObjectID(), Name: reqBody.Name, Users: []models.User {user}}
	createNew, err := projectCollection.InsertOne(context.TODO(), newProject)

	if err != nil {
		c.JSON(http.StatusBadRequest,  gin.H{"error: ": err.Error()})
	}

	update := bson.D{
		{"$set", bson.D{
			{"projects", []models.Project{newProject}},
		}},
	}

	result, err := GetUserCollection().UpdateOne(context.TODO(), user, update)

	if err != nil {
		c.JSON(http.StatusBadRequest,  gin.H{"error: ": err.Error()})
	}
	fmt.Println(result)

	c.JSON(http.StatusOK, createNew)
}
