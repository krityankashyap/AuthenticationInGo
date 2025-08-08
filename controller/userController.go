package controller

import (
	"AuthInGo/services"
	"net/http"
)

type UserController struct { // this struct is gonna depend on the service
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
  return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	uc.UserService.CreateUser() // Here u typically parse the request body, validate input, and call the service to create user
	w.Write([]byte("User register endpoint"))
}