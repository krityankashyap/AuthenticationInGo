package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(planePassword string) (string , error) {
   hash,err := bcrypt.GenerateFromPassword([]byte(planePassword) , bcrypt.DefaultCost)

	 if err!= nil {
		fmt.Println("Error while hashing the password ", err)
		return "",err
	 }
	 return string(hash),nil
}

func CheckHashedPassword(planePassword string, hashedPassword string) bool { // this function is going to check wheather the user tries to login for the first time then the plane password is correct or not
 
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword) , []byte(planePassword))

	return  err == nil
}