package auth

import (
	"backend/loggers"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (s *service) GetAuthCallbackHandler() http.HandlerFunc {

	// TODO: make login error page(s)
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")

		// add provider to the existing context
		loggers.Debug.Printf("Adding provider %s to context", provider)
		newCtx := context.WithValue(r.Context(), "provider", provider)
		r = r.WithContext(newCtx)

		loggers.Debug.Println("Getting user from gothic...")
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			loggers.Error.Printf("Error completing auth: %v", err)
			http.Redirect(w, r, "/login-error", http.StatusSeeOther)
			return
		}

		// check if authenticated user is an admin in the database
		admin, err := s.db.GetAdminByEmail(user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				loggers.Debug.Printf("No admin found with email %v", user.Email)
				// dandle non-admin user, redirect appropriately
				http.Redirect(w, r, "/login-unauthorized", http.StatusSeeOther)
				return
			} else {
				loggers.Error.Printf("Error getting admin: %v", err)
				http.Redirect(w, r, "/login-error", http.StatusSeeOther)
				return
			}
		}

		loggers.Debug.Printf("Admin found: %v", admin)

		// retrieve or create session
		loggers.Debug.Println("Retrieving session...")
		session, err := s.store.Get(r, "session-name")
		if err != nil {
			loggers.Error.Printf("Error retrieving session: %v", err)
			http.Redirect(w, r, "/login-error", http.StatusSeeOther)
			return
		}

		// store user and admin data in session

		// google info
		session.Values["userID"] = user.UserID
		session.Values["email"] = user.Email
		session.Values["name"] = user.Name
		session.Values["avatar_url"] = user.AvatarURL

		// postgres info
		session.Values["adminID"] = admin.ID
		session.Values["adminPosition"] = admin.Position

		// save session
		if err := session.Save(r, w); err != nil {
			loggers.Error.Printf("Error saving session: %v", err)
			http.Redirect(w, r, "/login-error", http.StatusSeeOther)
			return
		}

		// redirect to admin dashboard or another appropriate page
		http.Redirect(w, r, "http://localhost:5173/admin/dashboard", http.StatusFound)
	}
}

func (s *service) LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// clearing OAuth data
		loggers.Debug.Println("Clearing OAuth data...")
		provider := chi.URLParam(r, "provider")
		r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
		gothic.Logout(w, r)

		// next, clear application session data
		session, err := s.store.Get(r, "session-name")
		if err == nil {
			loggers.Debug.Println("Clearing application session data...")

			// delete session data
			session.Values["userID"] = nil
			session.Values["name"] = nil
			session.Values["email"] = nil
			session.Values["avatar_url"] = nil

			session.Options.MaxAge = -1
			// save changes
			session.Save(r, w)
		}

		// redirect to homepage
		http.Redirect(w, r, "http://localhost:5173", http.StatusTemporaryRedirect)
	}
}

func (s *service) BeginAuthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")

		r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

		gothic.BeginAuthHandler(w, r)
	}
}

func (s *service) SessionInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if user is authenticated
		session, err := s.store.Get(r, "session-name")
		if err != nil || session.Values["userID"] == nil {
			// user is not authenticated
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"authenticated": false})
			return
		}

		// respond with user info
		userInfo := map[string]interface{}{
			"authenticated": true,
			"userID":        session.Values["userID"],
			"name":          session.Values["name"],
			"email":         session.Values["email"],
			"avatar_url":    session.Values["avatar_url"],

			// admin info
			"adminID":       session.Values["adminID"],
			"adminPosition": session.Values["adminPosition"],
		}
		json.NewEncoder(w).Encode(userInfo)
	}
}
