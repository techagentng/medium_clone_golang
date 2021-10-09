package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5430
	user = "decagon"
	dbname = "medium"

)

// Db Inject psql instance into model layer for easy testing


func ConnectDb () (*sql.DB, error) {
	var Db *sql.DB
	var err error
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s dbname = %s sslmode = disable", host, port, user, dbname)
	Db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	if err = Db.Ping(); err != nil {
		return nil, err
		Db.Close()
	}
	fmt.Println("psql connected")
	return Db, nil
}
