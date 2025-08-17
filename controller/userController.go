package controller

import (
	"AuthInGo/services"
	"fmt"
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

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserByID called in UserController")
	uc.UserService.GetUserByID() // Here u typically parse the request body, validate input, and call the service to create user
	w.Write([]byte("User fetching endpoint done"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUse called in UserController")
	uc.UserService.CreateUser() // Here u typically parse the request body, validate input, and call the service to create user
	w.Write([]byte("User fetching endpoint done"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login User called in LoginUser controller")
	uc.UserService.LoginUser()
	w.Write([]byte("User fetching endpoint done"))
}