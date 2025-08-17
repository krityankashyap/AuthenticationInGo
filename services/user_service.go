package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetUserByID() error
	CreateUser()  error
	LoginUser()   error
}

type UserServiceImp struct {
    UserRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImp{
		UserRepository: _userRepository,
	}
}

func (u *UserServiceImp) GetUserByID() error{
	fmt.Println("Fetching user in userService")
	u.UserRepository.GetByID() // this will call the repository to create the user
	return nil
}

func (u *UserServiceImp) CreateUser() error {
	fmt.Println("Creating User in UserService")
	password := "example_password"
	hashedPassword,err := utils.HashPassword(password)

	if err!=nil {
		return nil
	}

	u.UserRepository.Create("example_Name_1", "example1@gmail.com", hashedPassword)

	return nil
}

func (u *UserServiceImp) LoginUser() error {
	response := utils.CheckHashedPassword("example_password" , "$2a$10$xqV5vZE53sxWSRhAJyCGgewusfhygXMRR2LZB4q26e8MGJKwMA.hO")
  
	fmt.Println("Login response: ", response)

	return nil

}