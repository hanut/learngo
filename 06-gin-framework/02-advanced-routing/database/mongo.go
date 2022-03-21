package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// DefaultMongoUri is the default uri that is used incase not sent in via the environment
	// variable MONGO_URI
	DefaultMongoUri = "mongodb://admin:qweasd123@localhost:27017/"
)

var (
	MongoClient *mongo.Client
	MongoDb     *mongo.Database
	ColUser     *mongo.Collection
)

// Connect connects to mongodb
func InitMongo() error {
	uri := os.Getenv("CA_MONGO_URI")

	if len(uri) == 0 {
		uri = DefaultMongoUri
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Ping the mongo server to check the connection works
	pctx, pcanc := context.WithTimeout(context.Background(), 20*time.Second)
	defer pcanc()
	err = client.Ping(pctx, nil)

	if err != nil {
		return err
	}

	log.Println("Mongo DB connected successfully...")

	MongoClient = client
	MongoDb = client.Database("crudapp")
	ColUser = MongoDb.Collection("users")

	return nil
}

func CloseMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := MongoClient.Disconnect(ctx)

	if err != nil {
		return err
	}

	return nil
}
