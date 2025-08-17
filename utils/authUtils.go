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