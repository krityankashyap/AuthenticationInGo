package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controller"
	// db "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
	dbConfig "AuthInGo/config/db"
	repo "AuthInGo/db/repositories"
)

// Config holds the configuration for the server.
type Config struct {
	Addr string // PORT
}

type Application struct {
	Config Config
	//Store  db.Storage // Assuming storage is defined in db/repositories/storage.go
}

// Constructor for Config
func NewConfig() Config {

	port  := config.Getstring("PORT" , ":8080")
	return Config{
		Addr: port,
	}
}

// Constructor for Application
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
		// Store: *db.NewStorage(),
	}
}

func (app *Application) Run() error {

	db,err :=dbConfig.SetupDB()
   if err != nil{
		fmt.Println("Error setting up database: ", err)
		return err
	 }

	ur := repo.NewUserRepository(db)
  us := services.NewUserService(ur)
	uc := controller.NewUserController(us)
  uRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter), // use the router setup function to handle requests 
		ReadTimeout:  10 * time.Second, // Set read timeout to 10 seconds
		WriteTimeout: 10 * time.Second, // Set write timeout to 10 seconds
	}

	fmt.Println("Starting server on", app.Config.Addr)

	return server.ListenAndServe()
}