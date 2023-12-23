package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"backend/internal/database"
	"backend/loggers"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {

	loggers.Info.Println("Starting server...")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		// no port provided, default to 8080
		loggers.Info.Println("No port provided, defaulting to 8080")
		port = 8080
	}

	loggers.Info.Println("Connecting to the database...")
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	loggers.Info.Println("Registering routes...")
	// declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	loggers.Info.Println("Server is ready to handle requests at", server.Addr)
	return server
}
