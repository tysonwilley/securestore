package models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

func connectToDB() *sql.DB {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "root:nosyt13@/secureStore")

	if err != nil {
		log.Panic(err.Error())
	}

	return db
}
