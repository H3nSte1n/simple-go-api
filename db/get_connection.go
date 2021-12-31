package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s:27017"
	host                 		 = "host.docker.internal"
)

// GetConnection - Retrieves a client to the DocumentDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, host)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}
	
	return client, ctx, cancel
}