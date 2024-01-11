package main

import (
	"backend/internal/server"
	"backend/loggers"
)

func main() {

	// initialize the server
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		loggers.Error.Fatalf("Error starting server: %s", err)
	}
}
