package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"hats-for-parties/config"
)

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

func CreateDbConnection() *mongo.Database {
	log.Printf("Creating database connection to database %s\n", config.ServiceConfig.DBName)
	client := connectToClient()

	db := client.Database(config.ServiceConfig.DBName)
	log.Println("Successfully created database connection")
	return db
}
