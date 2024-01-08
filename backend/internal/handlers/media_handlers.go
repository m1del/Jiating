package handlers

import (
	"backend/internal/s3service"
	"backend/loggers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HandlerDependencies struct {
	S3Service s3service.Service
}

func (deps *HandlerDependencies) GetYearsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		years, err := deps.S3Service.GetYears()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			loggers.Error.Printf("Error getting years: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(years)
	}
}

func (deps *HandlerDependencies) GetEventsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		year := chi.URLParam(r, "year")
		loggers.Debug.Printf("year: %v", year)
		events, err := deps.S3Service.GetEvents(year)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			loggers.Error.Printf("Error getting events: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	}
}

func (deps *HandlerDependencies) ListPhotosHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		year := chi.URLParam(r, "year")
		event := chi.URLParam(r, "event")
		loggers.Debug.Printf("year: %v, event: %v", year, event)
		photos, err := deps.S3Service.ListPhotos(year, event)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			loggers.Error.Printf("Error getting photos: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(photos)
	}
}

func (deps *HandlerDependencies) GetPhotosHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		year := chi.URLParam(r, "year")
		event := chi.URLParam(r, "event")
		loggers.Debug.Printf("year: %v, event: %v", year, event)
		photos, err := deps.S3Service.GetPhotos(r.Context(), year, event)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			loggers.Error.Printf("Error getting photos: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(photos)
	}
}
