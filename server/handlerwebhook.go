package server

import (
	"fmt"
	"net/http"
)

func (s *Server) HandleWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Webhook received")

	}
}
