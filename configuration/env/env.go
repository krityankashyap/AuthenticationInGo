package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)
func load() { // this function is going to techinally load dotenv file which should be present at root folder

  err := godotenv.Load() // this returns a error object that if there is any error it gonna throw that error

	if err != nil {
		// Log the error if the .env file isn't found or cannot loaded
		fmt.Println("Error loading .env file")
	}
}

func Getstring(key string , fallback string) string { // whenever we want to load the string value whose key is known to us we gonna call this function
  load() // Ensure that .env file is loaded before accessing the key

	value,response :=os.LookupEnv(key) // LookupEnv retrieves the value of the environment variable named by the key. If the variable is present in the environment the value (which may be empty) is returned and the boolean is true. Otherwise the returned value will be empty and the boolean will be false.

	if !response { // in case the boolean response is false i.e, we r not able to find any corresponding key value pair the key in lookup function
    return fallback
  }
	// if response is true
	return value
}

func Getint(key string , fallback int) int {
	load()

	value,response :=os.LookupEnv(key)

	if !response {
		return fallback
	}
	intValue,err :=strconv.Atoi(value)

	if err!= nil{ // there is some error in conversion
		fmt.Printf("Error converting %s to int: %v\n", key ,err)
    return fallback
	}
	return intValue
}

func Getbool(key string , fallback bool) bool {
	load()

	value , response := os.LookupEnv(key)

	if !response {
		return fallback
	}

	boolValue , err := strconv.ParseBool(value)
	if err!= nil{
		fmt.Println("error on converting string to bool" , err)
		return fallback
	}
	return boolValue

}
