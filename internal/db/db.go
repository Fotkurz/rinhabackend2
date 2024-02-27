package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var CONN *sql.DB

func Connect() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", port, user, pass, name,
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to db: %v", err))
	}

	CONN = conn
}
