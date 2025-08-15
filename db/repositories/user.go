package db

import (
	// "database/sql"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImple struct {
   db *sql.DB
}

func NewUserRepository(_db *sql.DB) *UserRepositoryImple {
	return &UserRepositoryImple{
       db: _db,
	}
}

func (u *UserRepositoryImple) Create() error {
	fmt.Println("User is created in UserRepository")
	return nil
} 

