package server

import (
	"backend/internal/auth"
	"backend/internal/handlers"
	"backend/internal/s3service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"golang.org/x/time/rate"
)

func (s *Server) RegisterRoutes() http.Handler {
	// enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // react dev server
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	// configure handler dependencies
	s3service := s3service.NewService(s.s3Client)
	deps := &handlers.HandlerDependencies{
		S3Service: s3service,
	}

	// initalize chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// public routes
	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)

	r.Get("/auth/{provider}/callback", handlers.GetAuthCallbackHandler())
	r.Get("/logout/{provider}", handlers.LogoutHandler())
	r.Get("/auth/{provider}", handlers.BeginAuthHandler())

	//media
	r.Get("/api/get/photoshoot-years", deps.GetYearsHandler())
	r.Get("/api/get/photoshoot-events/{year}", deps.GetEventsHandler())

	// email
	r.With(RateLimitMiddleware).Post("/api/send-email", handlers.ContactFormSubmissionHandler())

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

var limiter = rate.NewLimiter(1, 2)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
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
