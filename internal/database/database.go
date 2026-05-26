package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	db, err := sql.Open(
		"sqlite",
		"./storage/app.db",
	)

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	createUsersTable()
}

func createUsersTable() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT
			email VARCHAR(255) UNIQUE
			password TEXT
		)
	`

	_, err := DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}
