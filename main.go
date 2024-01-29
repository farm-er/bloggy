package main

import (
	"bloggy_api/routes"
)

func main() {

	server := routes.NewApiServer(":3000")
	server.Run()
}
