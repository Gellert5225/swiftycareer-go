package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var Database *mongo.Database

func ConnectToMongo() {
	uri := "mongodb://swiftycareer_dev:swiftycareer@23.242.50.55:27017/dev?authSource=admin"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	Database = Client.Database("dev")

	fmt.Println("Connected to MongoDB at 27017")
}
