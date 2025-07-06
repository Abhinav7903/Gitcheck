package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

// ServeHTTP implements http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// RunServer starts the HTTP server
func RunServer() {
	server := &Server{
		router: mux.NewRouter(),
	}
	server.RegisterRoutes()

	listenPort := ":8080"
	fmt.Printf("Server started on %s\n", listenPort)

	err := http.ListenAndServe(listenPort, server)
	if err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
