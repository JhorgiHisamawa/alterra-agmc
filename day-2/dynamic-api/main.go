package main

import (
	"api/config"
	"api/routers"
)

func main() {
	// start the server, and log if it fails
	e := routers.New()

	// connect to database
	config.ConnectDB()

	e.Logger.Fatal(e.Start(":8080"))
}
