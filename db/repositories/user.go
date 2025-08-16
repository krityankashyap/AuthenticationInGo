package db

import (
	// "database/sql"
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID() (*models.User,error)
	Create()  (error)
}

type UserRepositoryImple struct {
   db *sql.DB
}

func NewUserRepository(_db *sql.DB) *UserRepositoryImple {
	return &UserRepositoryImple{
       db: _db,
	}
}

func (u *UserRepositoryImple) Create() (error){
	 
		query := "INSERT INTO users (username,email,password) VALUES (?,?,?)"

		result,err := u.db.Exec(query, "krityan", "krityan@example.com" , "12345") // Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.

		if err != nil{
			fmt.Println("error while creating the user" , err)
			return err
		}

		rowAffected,rowerr := result.RowsAffected()

		if rowerr != nil{
			fmt.Println("Error in row Affected", rowerr)
			return rowerr
		}

		if rowAffected == 0{
			fmt.Println("User is not created , No rows are affected")
			return nil
		}

		fmt.Println("Creating user in repository", rowAffected) 
    return nil
}


func (u *UserRepositoryImple) GetByID() (*models.User , error) {
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
			return nil,err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil,err
		}
  }
  // step 4:- print the user details
	fmt.Println("User fetched successfully:", User)

	return nil,nil
} 

