package controller

import (
	dtos "AuthInGo/dtos"
	"AuthInGo/services"
	"AuthInGo/utils"
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

	var payload dtos.UserLogindto // this is the payload we gonna create

	err := utils.Readjson(*r, &payload)

	if err != nil{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while logging", err)
		return
	}

	validationError := utils.Validator.Struct(&payload)

	if validationError != nil {
     w.Write([]byte("Invalid input data"))
		 return
	}

	jwtToken , err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login", err)
		return
	}
  
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User loggin successfully", jwtToken)
	}
