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

var Store *sessions.CookieStore

func Init() {
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

	// assign store to export
	Store = store

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, "http://localhost:3000/auth/google/callback",
			"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"),
	)
}

func AuthMiddleware(next http.Handler) http.Handler {
	loggers.Debug.Println("AuthMiddleware called")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggers.Debug.Println("Retrieving session...")
		session, err := Store.Get(r, "session-name")
		if err != nil || session.Values["userID"] == nil {
			loggers.Debug.Println("User is not logged in")

			// Check if the request is an AJAX request
			if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
				loggers.Debug.Println("AJAX request detected, sending unauthorized status")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// For non-AJAX requests, redirect to login
			loggers.Debug.Println("Redirecting to login page...")
			http.Redirect(w, r, "/auth/google", http.StatusSeeOther)
			return
		}
		// User is authenticated; proceed with the request
		next.ServeHTTP(w, r)
	})
}
