package handlers

import (
	"backend/internal/email"
	"encoding/json"
	"net/http"
)

type contactFormRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func ContactFormSubmissionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contactForm contactFormRequest
		if err := json.NewDecoder(r.Body).Decode(&contactForm); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if err := email.SendEmail(contactForm.Name, contactForm.Email, contactForm.Subject, contactForm.Message); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// success response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
	}
}
