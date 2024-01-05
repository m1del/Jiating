package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/loggers"
	"encoding/json"
	"net/http"
)

func AdminDashboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggers.Debug.Println("Retreiving users from context...")
		user := r.Context().Value("user")

		if user == nil {
			loggers.Debug.Println("User is not logged in")
			http.Error(w, "You must be logged in to view this page", http.StatusForbidden)
			return
		}
		loggers.Debug.Printf("User: %v\n", user)

	}
}

func ListAdminHandler(db database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		admins, err := db.GetAllAdmins()
		if err != nil {
			loggers.Error.Printf("Error getting admins: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonResp, err := json.Marshal(admins)
		if err != nil {
			loggers.Error.Printf("Error handling JSON marshal: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// CORS
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResp)
		if err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
		}
	}
}

func CreateAdminHandler(db database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var admin models.Admin
		if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
			loggers.Error.Printf("Error decoding admin: %v", err)
			http.Error(w, "Bad Request", http.StatusInternalServerError)
			return
		}

		if err := db.CreateAdmin(admin); err != nil {
			loggers.Error.Printf("Error creating admin: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Admin created successfully"))
	}
}
