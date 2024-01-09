package handlers

import (
	"backend/loggers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (deps *HandlerDependencies) GetPresignedUploadURLHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse event id and image file from request
		event := chi.URLParam(r, "event")
		file := chi.URLParam(r, "file")

		url, err := deps.S3Service.GeneratePresignedUploadURL(event, file, 900)
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

		url, err := deps.S3Service.DevGeneratePresignedUploadURL(event, file, 900)
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
