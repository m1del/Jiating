package auth

import (
	"backend/loggers"
	"net/http"
)

func (s *service) AuthMiddleware(next http.Handler) http.Handler {
	loggers.Debug.Println("AuthMiddleware called")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggers.Debug.Println("Retrieving session...")
		session, err := s.store.Get(r, "session-name")
		if err != nil || session.Values["userID"] == nil {
			loggers.Debug.Println("User not logged in")

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
