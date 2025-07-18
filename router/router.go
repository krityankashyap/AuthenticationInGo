package router

import (
	"AuthInGo/controller"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux { // what this function is gonna return a specify type of chi

  router := chi.NewRouter() // NewRouter returns a new Mux object that implements the Router interface.
  
	router.Get("/ping", controller.PingHandler)
	return router
}