package db

import (
	"database/sql"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImple struct {
  db *sql.DB
}

func (u *UserRepositoryImple) Create() error {
	return nil
} 

