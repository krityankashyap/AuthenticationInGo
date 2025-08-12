package config

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	env "AuthInGo/config/env"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
 
cfg := mysql.NewConfig()

cfg.User= env.Getstring("DB_USER" , "root")
cfg.Passwd= env.Getstring("DB_Password" , "root")
cfg.Net= env.Getstring("DB_NET", "tcp")
cfg.Addr= env.Getstring("DB_ADDR" , "127.0.0.1:3306")
cfg.DBName= env.Getstring("DBName" , "airbnb_dev")

fmt.Println("Connecting to database" , cfg.DBName , cfg.FormatDSN())

  db,err :=sql.Open("mysql" , cfg.FormatDSN()) // it takes all the config properties and setup these in format data source name(DSN)

	if err!=nil{
		fmt.Println("Error connecting to database" , err)
		return nil,err
	}
  fmt.Println("trying to connect with database....")

	PingErr := db.Ping()

	if PingErr != nil {
      fmt.Println("Error pinging to the database" , PingErr)
			return nil,PingErr
	}


  fmt.Println("Connected to the database successfully")

	return db,nil
}