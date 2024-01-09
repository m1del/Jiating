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
	// initalize chi router
	r := chi.NewRouter()

	// enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // react dev server
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	// apply CORS middleware globally
	r.Use(func(next http.Handler) http.Handler {
		return c.Handler(next)
	})
	r.Use(middleware.Logger)

	// configure handler dependencies
	s3service := s3service.NewService()
	deps := &handlers.HandlerDependencies{
		S3Service: s3service,
	}

	// public routes
	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)

	// authentication Routes
	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}/callback", handlers.GetAuthCallbackHandler())
		r.Get("/logout/{provider}", handlers.LogoutHandler())
		r.Get("/{provider}", handlers.BeginAuthHandler())
	})

	// api routes
	r.Route("/api", func(r chi.Router) {
		// admin routes
		r.Route("/admins", func(r chi.Router) {
		})

		// event routes
		r.Route("/events", func(r chi.Router) {
			r.Post("/upload/{event}/{file}", deps.DevGetPresignedUploadURLHandler())
		})

		// meida photo routes
		r.Route("/photoshoot", func(r chi.Router) {
			r.Get("/years", deps.GetYearsHandler())
			r.Get("/events/{year}", deps.GetEventsHandler())
			r.Get("/list/{year}/{event}", deps.ListPhotosHandler())
			r.Get("/photos/{year}/{event}", deps.GetPhotosHandler())
		})

		// email routes
		r.With(RateLimitMiddleware).Post("/send-email", handlers.ContactFormSubmissionHandler())

		// session related routes
		r.Get("/session-info", s.sessionInfoHandler)

	})

	return r
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
