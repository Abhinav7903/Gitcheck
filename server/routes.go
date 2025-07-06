package server

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/webhook", s.HandleWebhook()).Methods("POST")

}
