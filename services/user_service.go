package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	CreateUser() error
}

type UserServiceImp struct {
    UserRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImp{
		UserRepository: _userRepository,
	}
}

func (u* UserServiceImp) CreateUser() error{
	fmt.Println("Creating user in userService")
	u.UserRepository.Create() // this will call the repository to create the user
	return nil
}