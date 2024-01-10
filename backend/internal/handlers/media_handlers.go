package handlers

import (
	"backend/loggers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (deps *HandlerDependencies) GetPhotoshootYearsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		years, err := deps.S3Service.GetPhotoshootYears(r.Context())
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			loggers.Error.Printf("Error getting years: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(years)
	}
}

func (deps *HandlerDependencies) GetPhotoshootEventsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		year := chi.URLParam(r, "year")
		loggers.Debug.Printf("year: %v", year)
		events, err := deps.S3Service.GetPhotoshootEvents(r.Context(), year)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			loggers.Error.Printf("Error getting events: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	}
}

func (deps *HandlerDependencies) ListPhotoshootPhotosHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		year := chi.URLParam(r, "year")
		event := chi.URLParam(r, "event")
		loggers.Debug.Printf("year: %v, event: %v", year, event)
		photos, err := deps.S3Service.ListPhotoshootPhotos(r.Context(), year, event)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			loggers.Error.Printf("Error getting photos: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(photos)
	}
}

func (deps *HandlerDependencies) GetPhotshootPhotosHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		year := chi.URLParam(r, "year")
		event := chi.URLParam(r, "event")
		loggers.Debug.Printf("year: %v, event: %v", year, event)
		photos, err := deps.S3Service.GetPhotoshootPhotos(r.Context(), year, event)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			loggers.Error.Printf("Error getting photos: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(photos)
	}
}
