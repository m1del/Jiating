package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/loggers"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func EventFormSubmitHandler(db database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			loggers.Error.Printf("Event json not parsed correctly: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		existingEvent, err := db.GetEventByID(event.ID)
		if err != nil && err != sql.ErrNoRows {
			loggers.Error.Printf("Error checking for existing event: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if existingEvent != nil {
			if err := db.UpdateEvent(event); err != nil {
				loggers.Error.Printf("Error updating event: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			if err := db.CreateEvent(event); err != nil {
				loggers.Error.Printf("Error creating event: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

		// success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Event added successfully"})
	}
}

func GetEventHandler(db database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the "id" query parameter from the request
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "ID parameter is required", http.StatusBadRequest)
			return
		}

		// Fetch the event by ID using the service method
		event, err := db.GetEventByID(id)
		if err != nil {
			loggers.Error.Printf("Error getting event: %v", err)
			http.Error(w, "Event not found", http.StatusNotFound)
			return
		}

		// Marshal the event to JSON
		jsonResp, err := json.Marshal(event)
		if err != nil {
			loggers.Error.Printf("Error handling JSON marshal: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set headers and write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResp)
		if err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
		}
	}
}

func (deps *HandlerDependencies) UploadEventImageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse event id and image file from request
		event := chi.URLParam(r, "event")
		file := chi.URLParam(r, "file")

		url, err := deps.S3Service.GenerateUploadURL(event, file, 900)
		if err != nil {
			loggers.Error.Printf("Error generating upload url: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// return presigned url
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"url": url})
	}
}

func (deps *HandlerDependencies) DevUploadEventImageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse event id and image file from request
		event := chi.URLParam(r, "event")
		file := chi.URLParam(r, "file")

		url, err := deps.S3Service.DevGenerateUploadURL(event, file, 900)
		if err != nil {
			loggers.Error.Printf("Error generating upload url: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// return presigned url
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"url": url})
	}
}
