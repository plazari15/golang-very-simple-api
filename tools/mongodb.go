package tools

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// var mongoContext *context.WithDeadline
// var mongoClient *mongo.Client

func CreateNewClient(database string) MongoInstance {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27045/"))
	mongoContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = mongoClient.Connect(mongoContext)
	db := mongoClient.Database(database)

	if err != nil {
		panic(err)
	}

	if err := mongoClient.Ping(mongoContext, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	mg := MongoInstance{
		Client: mongoClient,
		Db:     db,
	}

	return mg
}

// func disconnect() {
// 	if err := mongoClient.Disconnect(mongoContext); err != nil {
// 		panic(err)
// 	}
// }
