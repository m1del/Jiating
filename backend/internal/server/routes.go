package server

import (
	"backend/internal/auth"
	"backend/loggers"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth/gothic"
	"github.com/rs/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	// enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // react dev server
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	// initalize chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// public routes
	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)

	r.Get("/auth/{provider}/callback", s.getAuthCallbackHandler)
	r.Get("/logout/{provider}", s.logoutHandler)
	r.Get("/auth/{provider}", s.beginAuthHandler)

	// admin routes
	adminRouter := chi.NewRouter()
	adminRouter.Use(auth.AuthMiddleware)
	adminRouter.Get("/dashboard", s.adminDashboardHandler)

	// mount admin routes under /admin
	r.Mount("/admin", adminRouter)

	// route to get session info
	r.Get("/api/session-info", s.sessionInfoHandler)

	// apply cors middleware to all routes
	corsHandler := c.Handler(r)

	return corsHandler
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) getAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
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

// handle user logout
func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// clearing OAuth data
	loggesr.Debug.Println("Clearing OAuth data...")
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

// initiates the auth process
func (s *Server) beginAuthHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
}

// admin auth middleware
func (s *Server) adminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve the user from the context
	loggers.Debug.Println("Retrieving user from context...")
	user := r.Context().Value("user")

	// check if the user is actually set in the context
	if user == nil {
		// if no user is set, it means the user is not authenticated
		loggers.Debug.Println("User is not authenticated")
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}
	loggers.Debug.Printf("User: %v\n", user)

	// TODO: extracct info form user object if needed
	// e.g. popular user name, role, etc. to personalize the dashboard

	// TODO: generate and serve the admin dashboard page
}

func (s *Server) sessionInfoHandler(w http.ResponseWriter, r *http.Request) {
	// check if user is authenticated
	session, err := auth.Store.Get(r, "session-name")
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
	}
	json.NewEncoder(w).Encode(userInfo)
}
