package database

import (
	"BudgBackend/src/config"
	"database/sql"
	"log"
)

var dbUri string

func init() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dbUri = config.DBUri
}

func DBInit() *sql.DB {
	var err error
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		panic(err)
	}
	return db
}
