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
