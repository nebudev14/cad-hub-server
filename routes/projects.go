package routes

// import (
// 	"context"
// 	"github.com/NebuDev14/cad-hub-server/db"
// 	"github.com/NebuDev14/cad-hub-server/models"
// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"net/http"
// )

// var projectCollection *mongo.Collection

// type CreateProjectReq struct {
// 	UserId		string
// 	name			string
// }

// func StartProjects(group *gin.RouterGroup) {
// 	projectCollection = db.GetDB().Collection("projects")
// 	group.POST("/:id", CreateProject)
// }

// func CreateProject(c *gin.Context) {
// 	var reqBody CreateProjectReq
// 	if err := c.Copy().ShouldBindJSON(&reqBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
// 	}


// 	// newProject := models.Project{Id: primitive.NewObjectID(), Name: reqBody.name, Users: []}

// }
