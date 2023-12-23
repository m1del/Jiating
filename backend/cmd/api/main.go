package main

import (
	"backend/internal/auth"
	"backend/internal/server"
	"backend/loggers"
)

func main() {

	// initialize the auth package
	loggers.Info.Println("Initializing auth package...")
	auth.Init()

	// initialize the server
	loggers.Info.Println("Starting server...")
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		loggers.Error.Fatalf("Error starting server: %s", err)
	}
}
