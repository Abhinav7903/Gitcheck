package server

import (
	"encoding/json"
	"io"
	"jit/factory"
	"net/http"
	"strings"
	"time"
)

type GitHubPushPayload struct {
	Ref    string `json:"ref"`
	Before string `json:"before"`
	After  string `json:"after"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
	HeadCommit struct {
		ID        string `json:"id"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"head_commit"`
}

func (s *Server) HandleWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload GitHubPushPayload

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := json.Unmarshal(body, &payload); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Get event type from header
		action := r.Header.Get("X-GitHub-Event")

		// Extract branch name
		branch := strings.TrimPrefix(payload.Ref, "refs/heads/")

		// Parse timestamp
		timestamp, err := time.Parse(time.RFC3339, payload.HeadCommit.Timestamp)
		if err != nil {
			timestamp = time.Now().UTC()
		}

		model := factory.Model{
			Request_id:  payload.HeadCommit.ID,
			Author:      payload.HeadCommit.Author.Name,
			Action:      action,
			From_branch: payload.Before,
			To_branch:   branch,
			Created_at:  timestamp.UTC().Format(time.RFC3339),
		}

		if err := s.mongodb.Create(&model); err != nil {
			http.Error(w, "Failed to store webhook", http.StatusInternalServerError)
			return
		}

		s.respond(w, &ResponseMsg{
			Message: "Webhook stored successfully",
			Data:    model,
		}, http.StatusOK, nil)
	}
}
