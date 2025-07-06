package server

import (
	"encoding/json"
	"jit/factory"
	"net/http"
	"time"
)

func (s *Server) HandleWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model factory.Model

		if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		model.Created_at = time.Now().UTC().Format(time.RFC3339) // Ensure timestamp

		err := s.mongodb.Create(&model)
		if err != nil {
			http.Error(w, "Failed to store webhook", http.StatusInternalServerError)
			return
		}

		s.respond(w, &ResponseMsg{
			Message: "Webhook stored successfully",
			Data:    model,
		}, http.StatusOK, nil)
	}
}
