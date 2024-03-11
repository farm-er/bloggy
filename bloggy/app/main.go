package main

import (
	"github.com/farm-er/bloggy/server"
)

func main() {
	ser := server.CreateServer(":4242")
	ser.Run()
}
