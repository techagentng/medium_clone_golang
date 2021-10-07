package database

import (
	"database/sql"
	"fmt"
)

const (
	host = "localhost"
	port = 5430
	user = "decagon"
	dbname = "medium"

)

// Db Inject database instance into model layer for easy testing


func ConnectDb () *sql.DB {
	var Db *sql.DB
	var err error
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s dbname = %s sslmode = disable", host, port, user, dbname)
	Db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected")
	return Db
}
