package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient holds the raw client instance. Currently unused but reserved for advanced access.
var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

func InitMongoWith(uri, dbName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("🌐 Connecting to MongoDB at %s ...", uri)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("❌ MongoDB connection error: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ MongoDB ping error: %v", err)
	}

	log.Printf("✅ Successfully connected to MongoDB — using database: %s", dbName)

	MongoClient = client
	MongoDatabase = client.Database(dbName)
}
