package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/loggers"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func AdminDashboardHandler() http.HandlerFunc {
	// TODO: currently unused consider removal
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

func CreateAdminHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var admin models.Admin
		err := json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			loggers.Error.Printf("Error decoding json body: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		admin.Name = strings.ToLower(admin.Name)
		admin.Email = strings.ToLower(admin.Email)
		admin.Status = strings.ToLower(admin.Status)
		admin.Position = strings.ToLower(admin.Position)

		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()

		adminID, err := s.CreateAdmin(ctx, admin)
		if err != nil {
			loggers.Error.Printf("creating admin: %v", err)
			if errors.Is(err, context.Canceled) {
				http.Error(w, "Request canceled", http.StatusRequestTimeout)
				return
			}
			if errors.Is(err, context.DeadlineExceeded) {
				http.Error(w, "Request timed out", http.StatusRequestTimeout)
				return
			}
			http.Error(w, err.Error(), determineEmailStatusCode(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Admin created successfully", "id": adminID})
	}
}

func GetAllAdminsHandler(fetchFunc database.AdminFetchFunc, s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		pageSizeStr := r.URL.Query().Get("pageSize")

		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 {
			pageSize = 10
		}
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()

		admins, err := fetchFunc(ctx, page, pageSize)
		if err != nil {
			loggers.Error.Printf("reading admins: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError) // TODO: write more descriptive error
			return
		}

		if len(admins) == 0 {
			loggers.Error.Printf("No admins found")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "No admins found"})
			return
		}

		totalCount, err := s.GetAdminCount(r.Context())
		if err != nil {
			loggers.Error.Printf("getting total admin count: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response := struct {
			Admins     []models.Admin `json:"admins"`
			TotalCount int            `json:"totalCount"`
		}{
			Admins:     admins,
			TotalCount: totalCount,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
		}
	}
}

func GetAdminHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "param") // either ID or email
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()

		fieldName := "email"
		if isValidUUID(param) {
			fieldName = "id"
		}
		admin, err := s.GetAdmin(ctx, fieldName, param)

		if err != nil {
			loggers.Error.Printf("Error getting admin: %v", err)
			var errMsg string
			if err.Error() == "admin not found" {
				errMsg = "Admin not found"
				w.WriteHeader(http.StatusNotFound)
			} else {
				errMsg = "Internal Server Error"
				w.WriteHeader(http.StatusInternalServerError)
			}
			http.Error(w, errMsg, 0) // 0 to avoid setting status code twice
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(admin); err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
		}
	}
}

func GetAdminCountHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()

		count, err := s.GetAdminCount(ctx)
		if err != nil {
			loggers.Error.Printf("getting admin count: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]int{"count": count}); err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
		}
	}
}

func UpdateAdminHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse admin id from url
		adminID := chi.URLParam(r, "id")
		loggers.Debug.Print("adminID: ", adminID)
		if adminID == "" {
			loggers.Error.Printf("Admin ID not provided")
			http.Error(w, "missing admin ID", http.StatusBadRequest)
			return
		}
		// decode request body
		var admin models.Admin
		err := json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		admin.ID = adminID

		// use database service to update admin
		err = s.UpdateAdmin(r.Context(), admin)
		if err != nil {
			if err.Error() == "email already exists" {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			} else {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Admin updated successfully"})
	}
}

func DeleteAdminHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//param := chi.URLParam(r, "param") // either ID or email
		//ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		//defer cancel()

		// TODO refactor this to use a single function
	}
}

// func AssociateAdminWithEventHandler(s database.Service) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// decode json  into EventAuthor struct
// 		var eventAuthor models.EventAuthor
// 		err := json.NewDecoder(r.Body).Decode(&eventAuthor)
// 		if err != nil {
// 			loggers.Error.Printf("Error decoding json body: %v", err)
// 			http.Error(w, "Bad Request", http.StatusBadRequest)
// 			return
// 		}

// 		// use database service to associate admin with event
// 		err = s.AssociateAdminWithEvent(eventAuthor.AdminID, eventAuthor.EventID)
// 		if err != nil {
// 			loggers.Error.Printf("Error associating admin with event: %v", err)
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

//			// success response
//			w.WriteHeader(http.StatusOK)
//			w.Write([]byte("Admin associated with event successfully"))
//		}
//	}

// func DeleteAdminByIDHandler(s database.Service) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// extract admin id from url
// 		adminID := chi.URLParam(r, "adminID")
// 		err := s.DeleteAdminByID(adminID)
// 		if err != nil {
// 			if err.Error() == "cannot delete a permanent admin" {
// 				http.Error(w, "Forbidden", http.StatusForbidden)
// 				return
// 			}

// 			loggers.Error.Printf("Error deleting admin: %v", err)
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

// 		// success response
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Admin deleted successfully"))
// 	}
// }

// func DeleteAdminByEmailHandler(s database.Service) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// extract admin email from url and decode
// 		adminEmail := chi.URLParam(r, "adminEmail")
// 		decodedEmail, err := url.QueryUnescape(adminEmail)
// 		if err != nil {
// 			loggers.Error.Printf("Error decoding admin email: %v", err)
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}
// 		err = s.DeleteAdminByEmail(decodedEmail)
// 		if err != nil {
// 			if err.Error() == "cannot delete a permanent admin" {
// 				http.Error(w, "Forbidden", http.StatusForbidden)
// 				return
// 			}

// 			loggers.Error.Printf("Error deleting admin: %v", err)
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

// 		// success response
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Admin deleted successfully"))
// 	}
// }
