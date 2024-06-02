package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	Port           = ":8080"
	MongoURL       = "mongodb://localhost:27040"
	DB             *mongo.Client
	TodoDB         *mongo.Database
	TodoCollection *mongo.Collection
	MongoCtx       = context.Background()
)

func ConnectDatabase() {
	fmt.Println("Connecting to MongoDB...")
	clientOptions := options.Client().ApplyURI(MongoURL)
	client, err := mongo.Connect(MongoCtx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	if err := client.Ping(MongoCtx, nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	DB = client
	TodoDB = DB.Database("todoApp")
	TodoCollection = TodoDB.Collection("todos")
	fmt.Println("Connected to MongoDB")
}
