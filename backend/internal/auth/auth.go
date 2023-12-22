package auth

import (
	"log"
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

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, "http://localhost:3000/auth/google/callback"),
	)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := gothic.CompleteUserAuth(w, r); err != nil {
			// user is not authenticated, redirect to login page
			http.Redirect(w, r, "/auth/google", http.StatusSeeOther)
			return
		}
		// user is authenticated; proceed with the request
		next.ServeHTTP(w, r)
	})
}
