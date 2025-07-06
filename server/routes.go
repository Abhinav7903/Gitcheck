package server

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/webhook", s.HandleWebhook()).Methods("POST")
	s.router.HandleFunc("/get", s.handleGetRequest()).Methods("GET")
}
