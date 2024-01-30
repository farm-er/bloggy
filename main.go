package main

import (
	"bloggy_api/database"
	"bloggy_api/routes"
	"os"
)

func main() {

	os.Setenv("URL", "postgres://oussama:sqloussama@localhost:5432/blog")
	database.Connect()
	server := routes.NewApiServer(":3000")
	server.Run()
}
