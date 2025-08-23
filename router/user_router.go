package router

import (
	"AuthInGo/controller"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controller.UserController
}

func NewUserRouter(_userController *controller.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {

	 r.Get("/profile" , ur.userController.GetUserById)
	 r.With(middlewares.UserCreateRequestValidator).Post("/signup" , ur.userController.CreateUser)
	 r.With(middlewares.UserLoginRequestValidator).Post("/Login" , ur.userController.LoginUser)
}