package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/loggers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAuthorsByEventID(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse event id from request
		eventID := chi.URLParam(r, "eventID")
		if eventID == "" {
			http.Error(w, "Event ID is required", http.StatusBadRequest)
			return
		}

		// call database service
		authors, err := s.GetAuthorsByEventID(eventID)
		if err != nil {
			loggers.Error.Printf("Error getting authors: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// convert authors to json
		jsonResp, err := json.Marshal(authors)
		if err != nil {
			loggers.Error.Printf("Error marshalling authors: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// return authors
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResp)
		if err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func CreateEventHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.CreateEventRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			loggers.Error.Printf("Error decoding request body: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// construct event model from the request data
		newEvent := models.Event{
			EventName:   req.EventName,
			Date:        req.Date,
			Description: req.Description,
			Content:     req.Content,
			IsDraft:     req.IsDraft,
			Images:      req.Images,
		}

		// call database service
		eventID, err := s.CreateEvent(newEvent, req.AuthorIDs)
		if err != nil {
			loggers.Error.Printf("Error creating event: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// success response, return event ID
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"id": eventID})
	}
}

func UpdateEventByIDHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract eventID
		eventID := chi.URLParam(r, "eventID")
		if eventID == "" {
			http.Error(w, "Event ID is required", http.StatusBadRequest)
			return
		}

		// decode request body
		var req models.UpdateEventRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			loggers.Error.Printf("Error decoding request body: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// call databse service
		if err := s.UpdateEventByID(eventID, req); err != nil {
			loggers.Error.Printf("Error updating event by ID: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Event updated successfully"))
	}
}

func GetEventByIDHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//parse event id from request
		eventID := chi.URLParam(r, "eventID")
		if eventID == "" {
			http.Error(w, "Event ID is required", http.StatusBadRequest)
			return
		}

		// call database service
		event, err := s.GetEventByID(eventID)
		if err != nil {
			loggers.Error.Printf("Error getting event: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// convert event to json
		jsonResp, err := json.Marshal(event)
		if err != nil {
			loggers.Error.Printf("Error marshalling event: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		// return event
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResp)
		if err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func GetLastSevenPublishedEventsHandler(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call database service
		events, err := s.GetLastSevenPublishedEvents()
		if err != nil {
			loggers.Error.Printf("Error getting events: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		// convert events to json
		jsonResp, err := json.Marshal(events)
		if err != nil {
			loggers.Error.Printf("Error marshalling events: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		// return events
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResp)
		if err != nil {
			loggers.Error.Printf("Error writing response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
}

func (deps *HandlerDependencies) GetPresignedUploadURLHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse event id and image file from request
		event := chi.URLParam(r, "event")
		file := chi.URLParam(r, "file")

		url, err := deps.S3Service.GenerateEventImageUploadURL(event, file, 900)
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

func (deps *HandlerDependencies) DevGetPresignedUploadURLHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse event id and image file from request
		event := chi.URLParam(r, "event")
		file := chi.URLParam(r, "file")

		url, err := deps.S3Service.DevGenerateEventImageUploadURL(event, file, 900)
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
