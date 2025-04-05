package connections

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Global variables to store the MongoDB client and database connection
var (
	mongoClient *mongo.Client
	mongoDB     *mongo.Database
	once        sync.Once
)

// Connect initializes the MongoDB connection once
func Connect() {
	once.Do(func() {
		mongoDB = connectToMongoDB()
	})
}

// GetDB returns the MongoDB database connection
func GetDB() *mongo.Database {
	if mongoDB == nil {
		Connect()
	}
	return mongoDB
}

// GetClient returns the MongoDB client
func GetClient() *mongo.Client {
	if mongoClient == nil {
		Connect()
	}
	return mongoClient
}

// connectToMongoDB establishes a connection to MongoDB
func connectToMongoDB() *mongo.Database {
	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}

	databaseName := os.Getenv("MONGODB_DATABASE_NAME")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	var err error
	mongoClient, err = mongo.Connect(opts)

	if err != nil {
		panic(err)
	}

	databaseConnection := mongoClient.Database(databaseName)
	var result bson.M
	if err := databaseConnection.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	log.Printf(`Connected to MongoDB "%s" database`, databaseName)

	return databaseConnection
}
