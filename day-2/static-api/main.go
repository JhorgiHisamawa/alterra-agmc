package main

import "api/routers"

func main() {
	// start the server, and log if it fails
	e := routers.New()
	e.Logger.Fatal(e.Start(":8080"))
}
