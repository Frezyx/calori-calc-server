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

	userAuthRoute := userRoute.PathPrefix("/private").Subrouter()
	userAuthRoute.Use(s.authenticateUser)

	userAuthRoute.HandleFunc("/me", s.handleGetUserNow()).Methods("GET")
	userAuthRoute.HandleFunc("/delete/{id}", s.handleDeleteUser()).Methods("DELETE")
	userAuthRoute.HandleFunc("/edit/{id}", s.handleEditUser()).Methods("PUT")
}
