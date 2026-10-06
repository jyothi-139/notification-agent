package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {

	log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("DB_PORT:", os.Getenv("DB_PORT"))
	log.Println("DB_NAME:", os.Getenv("DB_NAME"))
	log.Println("DB_USER:", os.Getenv("DB_USER"))

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	log.Println("Attempting PostgreSQL Connection")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Database Ping Error:", err)
		return nil, err
	}

	var dbName string

	err = db.QueryRow(
		"SELECT current_database()",
	).Scan(&dbName)

	if err != nil {
		return nil, err
	}

	log.Println("Connected Database:", dbName)

	return db, nil
}
