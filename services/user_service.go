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

	// Prerequisite : this function will be given email and password as parameter, which we can hardcode for now

	// step 1:- Make a repository call to get user by email

	// step 2:- if user exists or not. if not return error

	// step 3:- if user exists then check the password by utils.Checkhashedpassword

	// step 4:- if the password matches, print JWT token else return error saying password doesn't match

	return nil

}