package main

import (
	"AuthInGo/app"
	config "AuthInGo/configuration/env"
)

func main() {
   config.Load() // since we r calling Load function again and again before caling any env variable so its better to call once before server starts 

	cfg := app.NewConfig() // Set the server to listen on port 8080
	app := app.NewApplication(cfg)

	app.Run() // call the run function from app folder
}
