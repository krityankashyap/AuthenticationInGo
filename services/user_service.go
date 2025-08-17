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
     
	email := "example1@gmail.com"
	password := "example_password"
	// step 1:- Make a repository call to get user by email
      
	 user,err := u.UserRepository.GetUserByEmail(email)

	 if err != nil {
		fmt.Println("User not found")
		return err
	 }

	// step 2:- if user exists or not. if not return error

	if err == nil {
		fmt.Println("No user found with given email")
		return nil
	}
	// step 3:- if user exists then check the password by utils.Checkhashedpassword
   isPasswordValid := utils.CheckHashedPassword(password , user.Password)

	 if !isPasswordValid{
		fmt.Println("Not a valid password for provided email")
		return nil
	 }
	// step 4:- if the password matches, print JWT token else return error saying password doesn't match
 
	fmt.Println("User logged in successfully, JWT token will be generated here")
	return nil

}