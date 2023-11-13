package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToPostgres() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	connString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(fmt.Sprintf("db: failed to connect to database, %v", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("db: failed to ping database, %v", err))
	}

	return db
}
