package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitConnection() {
	var err error
	connectionString := "user=postgres dbname=demo"
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

}
