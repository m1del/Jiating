package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/loggers"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
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

func GetAllAdminsHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		admins, err := s.GetAllAdmins()
		if err != nil {
			loggers.Error.Printf("Error getting admins: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// convert from Go data structure to JSON
		jsonResp, err := json.Marshal(admins)
		if err != nil {
			loggers.Error.Printf("Error handling JSON marshal: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResp)
		if err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func GetAllAdminsExceptFounderHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		admins, err := s.GetAllAdminsExceptFounder()
		if err != nil {
			loggers.Error.Printf("Error getting admins: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// convert from Go data structure to JSON
		jsonResp, err := json.Marshal(admins)
		if err != nil {
			loggers.Error.Printf("Error handling JSON marshal: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResp)
		if err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func CreateAdminHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// decode json body into admin struct
		var admin models.Admin
		err := json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			loggers.Error.Printf("Error decoding json body: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// use database service to create admin
		adminID, err := s.CreateAdmin(admin)
		if err != nil {
			loggers.Error.Printf("Error creating admin: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// success response and return id
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"id": adminID})
	}
}

func AssociateAdminWithEventHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// decode json  into EventAuthor struct
		var eventAuthor models.EventAuthor
		err := json.NewDecoder(r.Body).Decode(&eventAuthor)
		if err != nil {
			loggers.Error.Printf("Error decoding json body: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// use database service to associate admin with event
		err = s.AssociateAdminWithEvent(eventAuthor.AdminID, eventAuthor.EventID)
		if err != nil {
			loggers.Error.Printf("Error associating admin with event: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Admin associated with event successfully"))
	}
}

func DeleteAdminByIDHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract admin id from url
		adminID := chi.URLParam(r, "adminID")
		err := s.DeleteAdminByID(adminID)
		if err != nil {
			if err.Error() == "cannot delete a permanent admin" {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			loggers.Error.Printf("Error deleting admin: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Admin deleted successfully"))
	}
}

func DeleteAdminByEmailHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract admin email from url and decode
		adminEmail := chi.URLParam(r, "adminEmail")
		decodedEmail, err := url.QueryUnescape(adminEmail)
		if err != nil {
			loggers.Error.Printf("Error decoding admin email: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = s.DeleteAdminByEmail(decodedEmail)
		if err != nil {
			if err.Error() == "cannot delete a permanent admin" {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			loggers.Error.Printf("Error deleting admin: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Admin deleted successfully"))
	}
}

func GetAdminByIDHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract admin id from url
		adminID := chi.URLParam(r, "adminID")
		admin, err := s.GetAdminByID(adminID)
		if err != nil {
			loggers.Error.Printf("Error getting admin: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// convert from Go data structure to JSON
		jsonResp, err := json.Marshal(admin)
		if err != nil {
			loggers.Error.Printf("Error handling JSON marshal: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResp)
		if err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func GetAdminByEmailHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract admin email from url
		adminEmail := chi.URLParam(r, "adminEmail")
		admin, err := s.GetAdminByEmail(adminEmail)
		if err != nil {
			loggers.Error.Printf("Error getting admin by email: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// convert from Go data structure to JSON
		jsonResp, err := json.Marshal(admin)
		if err != nil {
			loggers.Error.Printf("Error handling JSON marshal: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResp)
		if err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func UpdateAdminHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// decode json body into admin struct
		var admin models.Admin
		err := json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			loggers.Error.Printf("Error decoding json body: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// use database service to update admin
		err = s.UpdateAdmin(admin)
		if err != nil {
			loggers.Error.Printf("Error updating admin: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Admin updated successfully"))
	}
}
