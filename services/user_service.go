package services

import db "AuthInGo/db/repositories"

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
	return nil
}