package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func CreateDatabase() (*sql.DB, error) {

	connStr := "user=dev dbname=miniguide password=dev sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
