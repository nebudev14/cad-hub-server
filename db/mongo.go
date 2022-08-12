package db

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *mongo.Database
var mongoClient *mongo.Client

func InitConnection() {
	err := godotenv.Load("local.env")
	if err!= nil {
		fmt.Println(err)
	}

	conectionString := os.Getenv("CONNECTION_STRING")
	clientOptions := options.Client().ApplyURI(conectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil { panic(err) }

	err = client.Ping(context.TODO(), nil)
	if err != nil { panic(err) }

	mongoClient = client
	mongoDB = client.Database("cads")
}


