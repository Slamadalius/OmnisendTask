package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/BurntSushi/toml"
)

var DB *mongo.Database
var CTX context.Context

// Represents database server and credentials
type DbConfig struct {
	ConnString string
	DbName     string
}

func configDB(ctx context.Context) (*mongo.Database, error) {
	var config DbConfig

	// Read and parse the configuration file
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	// Creating new mongo client and passing connection string
	client, err := mongo.NewClient(options.Client().ApplyURI(config.ConnString))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	reviewsDb := client.Database(config.DbName)

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