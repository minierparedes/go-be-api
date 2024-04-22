package db

import (
	"database/sql"
	"log"
)

func NewPostgresStorage(cfg string) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
