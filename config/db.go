package config

import (
	"fmt"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var CTX context.Context

func configDB(ctx context.Context) (*mongo.Database, error) {
	// Creating new mongo client and passing connection string (getting string from env variable)
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_STRING")))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDb")

	// Using database set in env variable
	reviewsDb := client.Database(os.Getenv("DATABASE_NAME"))

	// Returning pointer to a database
	return reviewsDb, nil
}

func init() {
	var err error
	// The mongo.Database initialization process requires a context.Context object
	CTX = context.Background()
	CTX, cancel := context.WithCancel(CTX)
	defer cancel()

	// Calling DB connection function giving context object
	DB, err = configDB(CTX)
	if err != nil {
		log.Fatal(err)
	}
}