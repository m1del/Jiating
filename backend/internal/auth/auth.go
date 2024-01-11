package auth

import (
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
}

func NewService(store *sessions.CookieStore) Service {
	return &service{
		store: store,
	}
}

func NewAuth() Service {
	err := godotenv.Load()
	if err != nil {
		loggers.Error.Fatal("Error loading .env file")
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if googleClientID == "" || googleClientSecret == "" {
		loggers.Error.Fatal("Missing Google Client ID or Client Secret")
	}

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, "http://localhost:3000/auth/google/callback",
			"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"),
	)

	service := NewService(store)
	return service
}
