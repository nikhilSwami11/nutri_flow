package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	log.Println("database connected successfully")
	return db
}
