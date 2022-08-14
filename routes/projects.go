package routes

import (
	"context"

	"github.com/NebuDev14/cad-hub-server/db"
	"github.com/NebuDev14/cad-hub-server/db/helpers"
	"github.com/NebuDev14/cad-hub-server/models"
	"github.com/gin-gonic/gin"

	// "go.mongodb.org/mongo-driver/bson"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var projectCollection *mongo.Collection = db.GetDB().Collection("projects")

type CreateProjectReq struct {
	UserId		string
	name			string
}

func StartProjects(group *gin.RouterGroup) {
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

	newProject := models.Project{Id: primitive.NewObjectID(), Name: reqBody.name, Users: []models.User {user}}
	createNew, err := projectCollection.InsertOne(context.TODO(), newProject)

	if err != nil {
		c.JSON(http.StatusBadRequest,  gin.H{"error: ": err.Error()})
	}

	c.JSON(http.StatusOK, createNew)
}
