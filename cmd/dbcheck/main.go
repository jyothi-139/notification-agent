package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"notification-consumer/pkg/postgres"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	db, err := postgres.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query(`
		SELECT id, email, subject, status
		FROM notifications
		ORDER BY id DESC
		LIMIT 10
	`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("Recent Notifications:")
	fmt.Println("---------------------")

	for rows.Next() {

		var id int
		var email string
		var subject string
		var status string

		if err := rows.Scan(
			&id,
			&email,
			&subject,
			&status,
		); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"ID=%d | Email=%s | Subject=%s | Status=%s\n",
			id,
			email,
			subject,
			status,
		)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
