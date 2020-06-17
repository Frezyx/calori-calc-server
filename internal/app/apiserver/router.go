package apiserver

import "github.com/gorilla/handlers"

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	// /user
	userRoute := s.router.PathPrefix("/user").Subrouter()
	userRoute.HandleFunc("/create", s.handleUsersCreate()).Methods("POST")
	userRoute.HandleFunc("/auth", s.handleSessionsCreate()).Methods("POST")

	// После регистарции
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)

	private.HandleFunc("/me", s.handleGetUserNow()).Methods("GET")
}
