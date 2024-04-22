package main

import (
	"cms-app/cmd/api"
	"cms-app/config"
	"cms-app/db"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Env.DBHost, config.Env.DBPort, config.Env.DBUser, config.Env.DBPassword, config.Env.DBName, config.Env.DBSSLMode)

	db, err := db.NewPostgresStorage(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to storage")
}
