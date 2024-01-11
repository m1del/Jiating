package auth

import (
	"backend/internal/database"
	"backend/loggers"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "random string for demo purposes only"
	MaxAge = 86400 * 30 // 30 days
	IsProd = false
)

type AuthConfig struct {
	Store        *sessions.CookieStore
	DB           database.Service
	ClientID     string
	ClientSecret string
}

// SOLID: Interface Segregation Principle :)
type Service interface {
	GetAuthCallbackHandler() http.HandlerFunc
	LogoutHandler() http.HandlerFunc
	BeginAuthHandler() http.HandlerFunc
	AuthMiddleware(next http.Handler) http.Handler
	SessionInfoHandler() http.HandlerFunc
}

type service struct {
	store *sessions.CookieStore
	db    database.Service
}

func NewAuth(config *AuthConfig) Service {
	service := &service{
		store: config.Store,
		db:    config.DB,
	}

	setUpGoth(config.ClientID, config.ClientSecret)
	return service

}

func setUpGoth(clientID, clientSecret string) {
	goth.UseProviders(
		google.New(
			clientID, clientSecret, "http://localhost:3000/auth/google/callback",
			"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile",
		),
	)
}

func LoadAuthConfig(db database.Service) (*AuthConfig, error) {
	loggers.Debug.Println("Loading auth config...")
	err := godotenv.Load()
	if err != nil {
		loggers.Error.Fatal("Error loading .env file")
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if googleClientID == "" || googleClientSecret == "" {
		loggers.Error.Fatal("Missing Google Client ID or Client Secret")
	}

	// setup cookie store
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	// setup cookie options
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = IsProd

	gothic.Store = store

	return &AuthConfig{
		Store:        store,
		DB:           db,
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
	}, nil
}
