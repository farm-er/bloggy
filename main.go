package main

import (
	"bloggy_api/database"
	"bloggy_api/routes"
	"context"
	"os"
)

func main() {

	os.Setenv("POSTURL", "postgres://oussama:sqloussama@localhost:5432/blog")
	os.Setenv("MONGOURL", "mongodb://localhost:27017")
	database.PostConnect()
	client := database.MongoConnect()

	server := routes.NewApiServer(":3000", client)
	server.Run()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
