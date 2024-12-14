package models

import (
	"context"
	"fmt"
	"proxy/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(data *types.User) error {
	var dbName = "DB"
	var collectionName = "users"
	collection := DB.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	newUser := types.User{
		PrivateKey: data.PrivateKey,
		Publickey:  data.Publickey,
		Hash:       data.Hash,
	}
	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		fmt.Println("Error saving user", err)
		return err
	}
	err = CloseDB()
	if err != nil {
		fmt.Println("Error closing DB", err)
	}
	fmt.Println("User Created", result)
	return nil
}

func GetUser(publicKey string) (types.User, error) {
	var dbName = "DB"
	var collectionName = "users"
	collection := DB.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result types.User
	var filter = bson.M{"publickey": publicKey}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("Error fetching user", err)
		return types.User{}, err
	}
	return result, nil
}
