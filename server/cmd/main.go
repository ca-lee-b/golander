package main

import (
	"os"

	"github.com/ca-lee-b/golander/server/internal/api"
	"github.com/ca-lee-b/golander/server/internal/db"
	"github.com/ca-lee-b/golander/server/internal/log"
	"github.com/ca-lee-b/golander/server/internal/repositories"
	"github.com/joho/godotenv"
)

func main() {
	logger := log.New()
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	db := db.ConnectToPostgres()
	logger.Info("Connected to database")

	repositories := repositories.New(db)
	api := api.New(os.Getenv("PORT"), repositories, logger)

	err = api.Listen()
	if err != nil {
		logger.Error("Failed to gracefully shutdown server, %v", err)
	}

	err = db.Close()
	if err != nil {
		logger.Error("Failed to close database connection, %v", err)
	}

	logger.Info("Successfully shutdown")
}
