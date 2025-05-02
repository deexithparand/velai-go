package db

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoClient    *mongo.Client
	UserCollection *mongo.Collection
	TaskCollection *mongo.Collection
	Ctx            = context.TODO() // You can replace this with a custom context if needed
)

// ConnectMongo sets up the connection to MongoDB and initializes collections.
func ConnectMongo(uri string, dbName string) error {

	// Create a context with timeout to ensure no operation hangs indefinitely
	ctx, cancel := context.WithTimeout(Ctx, 10*time.Second)
	defer cancel()

	// Connect to MongoDB using the context
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Ping MongoDB to verify the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	log.Info("Pinged your deployment. You successfully connected to MongoDB!")

	// Initialize the MongoDB collections
	MongoClient = client
	UserCollection = client.Database(dbName).Collection("users")
	TaskCollection = client.Database(dbName).Collection("tasks")

	return nil
}
