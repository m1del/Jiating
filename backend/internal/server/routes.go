package server

import (
	"backend/internal/auth"
	"context"
	"encoding/json"
	"fmt"
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
	newCtx := context.WithValue(r.Context(), "provider", provider)
	r = r.WithContext(newCtx)

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		// log the error for internal tracking
		log.Printf("error completing auth: %v", err)
		// redirect the user to a login error page or display an error message
		http.Redirect(w, r, "/login-error", http.StatusSeeOther)
		return
	}

	// create session or retrive existing
	session, err := auth.Store.Get(r, "session-name")
	if err != nil {
		log.Printf("error retrieving session: %v", err)
		http.Redirect(w, r, "/login-error", http.StatusSeeOther)
		return
	}

	fmt.Println("user id: ", user.UserID)
	fmt.Println("email: ", user.Email)
	fmt.Println("name: ", user.Name)
	fmt.Println("avatar url: ", user.AvatarURL)
	fmt.Println("first name: ", user.FirstName)
	// store user data in session
	session.Values["userID"] = user.UserID
	session.Values["email"] = user.Email
	session.Values["name"] = user.Name
	session.Values["avatar_url"] = user.AvatarURL

	// save session
	if err := session.Save(r, w); err != nil {
		log.Printf("error saving session: %v", err)
		http.Redirect(w, r, "/login-error", http.StatusSeeOther)
		return
	}

	// redirect to a post-login page, such as the admin dashboard or home page
	http.Redirect(w, r, "http://localhost:5173/admin/dashboard", http.StatusFound)
}

// handle user logout
func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// clearing OAuth data
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	gothic.Logout(w, r)

	// next, clear application session data
	session, err := auth.Store.Get(r, "session-name")
	if err == nil {
		fmt.Println("clearing session data")
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
	user := r.Context().Value("user")

	// check if the user is actually set in the context
	if user == nil {
		// If no user is set, it means the user is not authenticated
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// TODO: extracct info form user object if needed
	// e.g. popular user name, role, etc. to personalize the dashboard

	// TODO: generate and serve the admin dashboard page

	// for a simple response for development purposes
	fmt.Fprintf(w, "Welcome to the Admin Dashboard")
}

func (s *Server) sessionInfoHandler(w http.ResponseWriter, r *http.Request) {
	// check if user is authenticated
	session, err := auth.Store.Get(r, "session-name")
	fmt.Printf("Session values: %v\n", session.Values)
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
