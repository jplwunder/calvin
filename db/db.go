package db

import (
	"context"
	"log"
	// "os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var collection *mongo.Collection

func ConnectToMongo() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.3.2")

	// username := os.Getenv("MONGO_DB_USERNAME")
	// password := os.Getenv("MONGO_DB_PASSWORD")

	// clientOptions.SetAuth(options.Credential{
	// 	Username: username,
	// 	Password: password,
	// })

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to mongo...")

	return client, nil
}
