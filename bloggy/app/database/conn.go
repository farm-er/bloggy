package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createClient() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	mongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoDb))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected")
	return client
}

var client *mongo.Client = createClient()

func OpenCollection(collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Bloggy").Collection(collectionName)
	return collection
}
