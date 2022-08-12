package db

import (

	"fmt"
	"os"

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// var mongoDB *mongo.Database
// var mongoClient *mongo.Client

func InitConnection() {
	err := godotenv.Load("local.env")
	if err!= nil {
		fmt.Println(err)
	}

	conectionString := os.Getenv("CONNECTION_STRING")
	fmt.Println(conectionString)
}
