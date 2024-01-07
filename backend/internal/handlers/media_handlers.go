package handlers

import (
	"backend/internal/s3service"
	"backend/loggers"
	"encoding/json"
	"net/http"
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
