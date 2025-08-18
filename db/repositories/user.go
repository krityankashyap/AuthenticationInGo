package db

import (
	// "database/sql"
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID() (*models.User,error)
	Create(usename string, email string, hashedPassword string)  (error)
	GetAll()  ([]*models.User,error)
	DeleteById(id int64)  error
	GetUserByEmail(email string)  (*models.User,error)

}

type UserRepositoryImple struct {
   db *sql.DB
}

func NewUserRepository(_db *sql.DB) *UserRepositoryImple {
	return &UserRepositoryImple{
       db: _db,
	}
}

func (u *UserRepositoryImple) Create(usename string, email string, hashedPassword string) (error){
	 
		query := "INSERT INTO users (username,email,password) VALUES (?,?,?)"

		result,err := u.db.Exec(query, usename, email , hashedPassword) // Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.

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

func (u *UserRepositoryImple) GetAll() ([]*models.User,error){
	return nil,nil
}

func (u *UserRepositoryImple) DeleteById(id int64) error{

	query := "DELETE FROM users WHERE id=?"

	result,err := u.db.Exec(query,id) // Exec executes a query without returning any rows

	if err != nil{
		fmt.Println("Error on deleting the user:",err)
		return err
	}

	affectedRow, errorRow := result.RowsAffected() // RowsAffected returns the number of rows affected by an update, insert, or delete.
 
	if errorRow != nil{
     fmt.Println("Error getting updated after deletion", errorRow)
		 return errorRow
	}

	if affectedRow == 0 {
    fmt.Println("No row is deleted, No rows r updated")
		return nil
	}
   fmt.Println("Deleting User in repository:", affectedRow)
	return nil
}

func (u *UserRepositoryImple) GetUserByEmail(email string) (*models.User, error) {
	fmt.Println("Fetching user by email")

	query := "SELECT id,email,password FROM users WHERE email = ?"

	
	row := u.db.QueryRow(query, email)
	
	users := &models.User{}

	err := row.Scan(&users.Id , &users.Email , &users.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user with given email")
			return nil,err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil,err
		}
	}
    return users,nil
}



