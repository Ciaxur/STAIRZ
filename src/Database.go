package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DatabaseOptions \
//  - Required Options for Database Connection
type DatabaseOptions struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

// InitDatabase \
//  - Initializes Database Connetion given Database Options
func InitDatabase(options DatabaseOptions) *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", options.host, options.port, options.user, options.password, options.dbname)
	log.Printf("Connecting to host %s\n", options.host)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err, "Open Database Error")

	// check db
	err = db.Ping()
	CheckError(err, "Database Ping Error")

	fmt.Println("Connected!")

	return db
}

// CheckError \
// Checks for Error, printing given Summary and panicing
//  if specified to do so
func CheckError(err error, summary string, isPanic ...bool) {
	if err != nil {
		log.Printf("Error: %s. \n%v\n", summary, err)
		if len(isPanic) > 0 && isPanic[0] == true {
			panic(err)
		}
	}
}
