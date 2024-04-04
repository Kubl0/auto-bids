package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(
			os.Getenv("DB_PROFILES_HOST")).SetServerAPIOptions(serverAPI),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var DB = ConnectDB()

func GetCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	return collection
}
