package main

import (
	"bloggy_api/database"
	"bloggy_api/routes"
)

func main() {

	const url string = "postgres://postgres:postgres@localhost/blog"
	sqlbase := database.NewDatabaseConnection(url)
	sqlbase.Connect()

	server := routes.NewApiServer(":3000")
	server.Run()

}
