package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDB connects to MongoDB and sets up the DB reference
func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	// Set DB variable to access MongoDB
	DB = client.Database("MTBS")
	if DB == nil {
		log.Fatal("Failed to initialize database")
	}
	log.Println("MongoDB connected")

}
