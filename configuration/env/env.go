package config

import (
	"fmt"
	"os"

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
