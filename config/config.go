package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MyPort() (string,error) {
	port := os.Getenv("PORT")
	if port == ""{
		port = "5001"
	}
	return ":" + port, nil
}

func Connectdb() (*sql.DB, error){
	db, errdb := sql.Open("mysql", "root:@tcp(localhost:3306)/goapi")
	if errdb != nil {
		return nil, errdb
	}
	err := db.Ping()
	return db, err
}