package main

import (
	"os"

	"github.com/ca-lee-b/golander/server/internal/api"
	"github.com/ca-lee-b/golander/server/internal/log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	logger := log.New()

	port := os.Getenv("PORT")
	if port == "" {
		logger.Error("No port environment variable found")
	}

	api := api.New(port, logger)

	err = api.Listen()
	if err != nil {
		logger.Error("Failed to gracefully shutdown server")
	}

	logger.Info("Successfully shutdown")
}
