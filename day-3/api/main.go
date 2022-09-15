package main

import (
	"api/config"
	"api/routers"

	"github.com/go-playground/validator"
)

func main() {
	// start the server, and log if it fails
	e := routers.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}
	// connect to database
	config.ConnectDB()

	e.Logger.Fatal(e.Start(":8080"))
}
