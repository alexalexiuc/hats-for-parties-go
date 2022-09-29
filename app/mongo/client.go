package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"hats-for-parties/config"
)

type MongoClient struct {
	Client         *mongo.Client
	HatsCollection *mongo.Collection
	LockFlag       *mongo.Collection
}

var MongoDbConn MongoClient

func InitMongoClient() {
	MongoDbConn.Client = connectToClient()
	MongoDbConn.HatsCollection = MongoDbConn.Client.Database(config.ServiceConfig.DBName).Collection(config.ServiceConfig.HatsCollectionName)
	MongoDbConn.LockFlag = MongoDbConn.Client.Database(config.ServiceConfig.DBName).Collection(config.ServiceConfig.LockFlagCollectionName)
}

func CloseMongoClient() {
	MongoDbConn.Client.Disconnect(context.Background())
}

func connectToClient() *mongo.Client {
	log.Println("Connecting to MongoDB...")
	// Connect to the database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.ServiceConfig.MongoConnString))

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Successfully connected and pinged.")

	return client
}
