package server

import (
	"fmt"
	mongodb "jit/pkg/db"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router  *mux.Router
	mongodb mongodb.Repository
}

// ServeHTTP implements http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// RunServer starts the HTTP server
func RunServer() {

	dbstring := "mongodb+srv://git:N2Ud0dzyUK54D7Hq@cluster0.xcypdzo.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	// Initialize the MongoDB repository
	if dbstring == "" {
		fmt.Println("Database connection string is empty")
		return
	}
	// connect to MongoDB
	if err := mongodb.NewRepository(dbstring).Connect(); err != nil {
		logrus.Error("Failed to connect to MongoDB:", err)
		return
	}

	server := &Server{
		router:  mux.NewRouter(),
		mongodb: mongodb.NewRepository(dbstring),
	}
	server.RegisterRoutes()

	listenPort := ":8080"
	fmt.Printf("Server started on %s\n", listenPort)

	err := http.ListenAndServe(listenPort, server)
	if err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
