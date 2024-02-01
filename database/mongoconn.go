package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGOURL")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// creating the database
	collection := client.Database("bloggy").Collection("posts")

	collection.InsertOne(context.TODO(), BlogPost{
		Author:           "official bloogy",
		Title:            "Welcome to BLOGGY",
		Description:      "this is the first post in bloggy",
		Content:          "We are honored to introduce you the new blogs website where everyone can share their thoughts.",
		Tags:             []string{"first"},
		Likes:            0,
		NumberOfComments: 0,
		Comments:         nil,
	})
	return client
}
