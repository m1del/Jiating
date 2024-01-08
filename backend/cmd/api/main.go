package main

import (
	"backend/internal/auth"
	"backend/internal/server"
	"backend/loggers"

	"github.com/joho/godotenv"
)

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		loggers.Error.Fatalf("Error loading .env file: %s", err)
	}

	// initialize the auth package
	loggers.Info.Println("Initializing auth package...")
	auth.NewAuth()

	// initialize the server
	server := server.NewServer()

	err = server.ListenAndServe()
	if err != nil {
		loggers.Error.Fatalf("Error starting server: %s", err)
	}
}
