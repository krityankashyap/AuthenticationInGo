package router

import (
	"AuthInGo/controller"
	// "AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface{ // Router interface defines the methods for the Router

 Register(r chi.Router) 
}

func SetupRouter(UserRouter Router) *chi.Mux { // what this function is gonna return a specify type of chi

  chiRouter := chi.NewRouter() // NewRouter returns a new Mux object that implements the Router interface.

	chiRouter.Use(middleware.Logger)
  
	chiRouter.Get("/ping", controller.PingHandler)

	UserRouter.Register(chiRouter)

	return chiRouter
}