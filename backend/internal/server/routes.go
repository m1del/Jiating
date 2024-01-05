package server

import (
	"backend/internal/auth"
	"backend/internal/handlers"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	// enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // react dev server
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	// initalize chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// public routes
	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)

	r.Get("/auth/{provider}/callback", handlers.GetAuthCallbackHandler())
	r.Get("/logout/{provider}", handlers.LogoutHandler())
	r.Get("/auth/{provider}", handlers.BeginAuthHandler())

	// admin routes
	adminRouter := chi.NewRouter()
	adminRouter.Use(auth.AuthMiddleware)
	adminRouter.Get("/dashboard", handlers.AdminDashboardHandler())      // handles admin dashboard
	adminRouter.Get("/list", handlers.ListAdminHandler(s.db))            // handles admin list
	adminRouter.Post("/create-admin", handlers.CreateAdminHandler(s.db)) // handles admin creation

	// mount admin routes under /admin
	r.Mount("/admin", adminRouter)

	// route to get session info
	r.Get("/api/session-info", s.sessionInfoHandler)

	// database routes

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

// func (s *Server) getAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
// 	provider := chi.URLParam(r, "provider")

// 	// add provider to the existing context, not overwrite it
// 	loggers.Debug.Printf("Adding provider %s to context", provider)
// 	newCtx := context.WithValue(r.Context(), "provider", provider)
// 	r = r.WithContext(newCtx)

// 	loggers.Debug.Println("Getting user from gothic...")
// 	user, err := gothic.CompleteUserAuth(w, r)
// 	if err != nil {
// 		// log the error for internal tracking
// 		loggers.Error.Printf("error completing auth: %v", err)
// 		// redirect the user to a login error page or display an error message
// 		http.Redirect(w, r, "/login-error", http.StatusSeeOther)
// 		return
// 	}

// 	// create session or retrive existing
// 	loggers.Debug.Println("Retreiving session...")
// 	session, err := auth.Store.Get(r, "session-name")
// 	if err != nil {
// 		loggers.Error.Printf("error retrieving session: %v", err)
// 		http.Redirect(w, r, "/login-error", http.StatusSeeOther)
// 		return
// 	}

// 	// store user data in session
// 	session.Values["userID"] = user.UserID
// 	session.Values["email"] = user.Email
// 	session.Values["name"] = user.Name
// 	session.Values["avatar_url"] = user.AvatarURL

// 	// save session
// 	if err := session.Save(r, w); err != nil {
// 		loggers.Error.Printf("error saving session: %v", err)
// 		http.Redirect(w, r, "/login-error", http.StatusSeeOther)
// 		return
// 	}

// 	// redirect to a post-login page, such as the admin dashboard or home page
// 	http.Redirect(w, r, "http://localhost:5173/admin/dashboard", http.StatusFound)
// }

// // handle user logout
// func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
// 	// clearing OAuth data
// 	loggers.Debug.Println("Clearing OAuth data...")
// 	provider := chi.URLParam(r, "provider")
// 	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
// 	gothic.Logout(w, r)

// 	// next, clear application session data
// 	session, err := auth.Store.Get(r, "session-name")
// 	if err == nil {
// 		loggers.Debug.Println("Clearing application session data...")

// 		// delete session data
// 		session.Values["userID"] = nil
// 		session.Values["name"] = nil
// 		session.Values["email"] = nil
// 		session.Values["avatar_url"] = nil

// 		session.Options.MaxAge = -1
// 		// save changes
// 		session.Save(r, w)
// 	}

// 	// redirect to homepage
// 	http.Redirect(w, r, "http://localhost:5173", http.StatusTemporaryRedirect)
// }

// // initiates the auth process
// func (s *Server) beginAuthHandler(w http.ResponseWriter, r *http.Request) {
// 	provider := chi.URLParam(r, "provider")

// 	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

// 	gothic.BeginAuthHandler(w, r)
// }

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
