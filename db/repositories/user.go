package db

import (
	// "database/sql"
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImple struct {
  // db *sql.DB
}

func NewUserRepository() *UserRepositoryImple {
	return &UserRepositoryImple{
      // db: db,
	}
}

func (u *UserRepositoryImple) Create() error {
	fmt.Println("User is created in UserRepository")
	return nil
} 

