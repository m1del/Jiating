package handlers

import (
	"encoding/json"
	"net/http"
)

type postFormRequest struct {
	Author    string `json:"author"`
	Event   string `json:"event"`
	Description string `json:"description"`
	Date string `json:"date"`
	Draft bool `json:"draft"`
}

func PostFormSubmitHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var postForm postFormRequest
		if err := json.NewDecoder(r.Body).Decode(&postForm); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
	}
}
