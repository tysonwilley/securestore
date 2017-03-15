package models

import (
	_ "github.com/go-sql-driver/mysql"
	"secureStore/config"
	"database/sql"
	"log"
	"fmt"
)

func connectToDB() *sql.DB {
	datasource := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s",
		config.Parameters.Database.User,
		config.Parameters.Database.Password,
		config.Parameters.Database.Host,
		config.Parameters.Database.Port,
		config.Parameters.Database.Database,
	)

	db, err := sql.Open("mysql", datasource)

	if err != nil {
		log.Panic(err.Error())
	}

	return db
}
