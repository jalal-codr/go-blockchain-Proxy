package models

import (
	"context"
	"fmt"
	"proxy/types"
	"time"
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
