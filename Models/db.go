package models

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func CloseDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := DB.Disconnect(ctx); err != nil {
		fmt.Println("Error disconnecting MongoDB:", err)
		return err
	}

	fmt.Println("Disconnected from MongoDB successfully.")
	return nil
}

func InitDb() {
	var mongoUri = os.Getenv("MONGO_URL")

	// CONNECT TO YOUR ATLAS CLUSTER:
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoUri,
	))

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("There was a problem connecting to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
		panic(err)
	}
	DB = client
	fmt.Println("Connected to MongoDB!")
}
