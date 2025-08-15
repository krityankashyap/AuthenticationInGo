package db

import (
	// "database/sql"
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID() error
}

type UserRepositoryImple struct {
   db *sql.DB
}

func NewUserRepository(_db *sql.DB) *UserRepositoryImple {
	return &UserRepositoryImple{
       db: _db,
	}
}

func (u *UserRepositoryImple) GetByID() error {
	fmt.Println("Fetching user by user repositories")

	// step 1:- prepare the query

	query := "SELECT id,username,email,password,created_at,updated_at FROM users WHERE id=?"

	// step 2:- execute the query
	row := u.db.QueryRow(query , 1)

	// step 3:- Process the result

	User := &models.User{}

  err := row.Scan(&User.Id, &User.Username, &User.Email, &User.Password, &User.CreatedAt, &User.UpdatedAt)    // Scan copies the columns from the matched row into the values pointed at by dest.

	if err != nil{
		if err == sql.ErrNoRows {
			fmt.Println("No user with given ID")
			return nil
		} else {
			fmt.Println("Error scanning user:", err)
			return err
		}
  }
  // step 4:- print the user details
	fmt.Println("User fetched successfully:", User)

	return nil
} 

