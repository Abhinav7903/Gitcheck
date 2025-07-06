package server

import (
	"encoding/json"
	"jit/factory"
	"net/http"
)

type ResponseMsg struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (s *Server) handleGetRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//queryParam := r.URL.Query().Get("param")
		requestId := r.URL.Query().Get("requestId")
		if requestId == "" {
			http.Error(w, "requestId query parameter is required", http.StatusBadRequest)
			return
		}

		// Process the request using the requestId
		var result *factory.Model

		result, err := s.mongodb.GetByRequestID(requestId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the result as JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

		//show the result in browser
		s.respond(w, result, http.StatusOK, nil)

	}
}
