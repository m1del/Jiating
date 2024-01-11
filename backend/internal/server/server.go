package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"backend/internal/auth"
	"backend/internal/database"
	"backend/internal/s3service"
	"backend/loggers"
	"backend/misc"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port     int
	db       database.Service
	s3Client *s3.Client
	auth     auth.Service
}

func NewServer() *http.Server {

	loggers.Info.Println("Starting server...")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		// no port provided, default to 8080
		loggers.Info.Println("No port provided, defaulting to 8080")
		port = 8080
	}

	// =========== Server setup =========== //

	loggers.Info.Println("Initializing S3 client...")
	s3Client, err := newS3Client()
	if err != nil {
		loggers.Error.Fatalf("failed to create S3 client: %v", err)
	}

	loggers.Info.Println("Initializing database...")
	dbClient := database.New(nil) // Error is handled in database.New

	loggers.Info.Println("Initializing auth service...")
	authConfig, err := auth.LoadAuthConfig(dbClient)
	if err != nil {
		loggers.Error.Fatalf("Error creating auth config: %v", err)
	}
	authService := auth.NewAuth(authConfig)

	loggers.Info.Println("Connecting to the database...")
	NewServer := &Server{
		port:     port,
		db:       dbClient,
		s3Client: s3Client,
		auth:     authService,
	}

	// seed the database
	if os.Getenv("ENV") == "dev" {
		// check if the database has already been seeded
		count, err := NewServer.db.GetAdminCount()
		if err != nil {
			loggers.Error.Fatalf("Error getting admin count: %v", err)
		}
		if count == 1 {
			loggers.Info.Println("Seeding the database...")
			err = misc.Seed(NewServer.db, 10, 30, 10)
			if err != nil {
				loggers.Error.Fatalf("Error seeding the database: %v", err)
			}
		}
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

func newS3Client() (*s3.Client, error) {
	cfg, err := s3service.NewAWSConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}
	return s3.NewFromConfig(cfg), nil
}
