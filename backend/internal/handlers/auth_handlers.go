package handlers

import (
	"backend/internal/auth"
	"backend/loggers"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func GetAuthCallbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")

		// add provider to the existing context, not overwrite it
		loggers.Debug.Printf("Adding provider %s to context", provider)
		newCtx := context.WithValue(r.Context(), "provider", provider)
		r = r.WithContext(newCtx)

		loggers.Debug.Println("Getting user from gothic...")
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			// log the error for internal tracking
			loggers.Error.Printf("error completing auth: %v", err)
			// redirect the user to a login error page or display an error message
			http.Redirect(w, r, "/login-error", http.StatusSeeOther)
			return
		}

		// create session or retrive existing
		loggers.Debug.Println("Retreiving session...")
		session, err := auth.Store.Get(r, "session-name")
		if err != nil {
			loggers.Error.Printf("error retrieving session: %v", err)
			http.Redirect(w, r, "/login-error", http.StatusSeeOther)
			return
		}

		// store user data in session
		session.Values["userID"] = user.UserID
		session.Values["email"] = user.Email
		session.Values["name"] = user.Name
		session.Values["avatar_url"] = user.AvatarURL

		// save session
		if err := session.Save(r, w); err != nil {
			loggers.Error.Printf("error saving session: %v", err)
			http.Redirect(w, r, "/login-error", http.StatusSeeOther)
			return
		}

		// redirect to a post-login page, such as the admin dashboard or home page
		http.Redirect(w, r, "http://localhost:5173/admin/dashboard", http.StatusFound)
	}
}

func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// clearing OAuth data
		loggers.Debug.Println("Clearing OAuth data...")
		provider := chi.URLParam(r, "provider")
		r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
		gothic.Logout(w, r)

		// next, clear application session data
		session, err := auth.Store.Get(r, "session-name")
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

func BeginAuthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")

		r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

		gothic.BeginAuthHandler(w, r)
	}
}
