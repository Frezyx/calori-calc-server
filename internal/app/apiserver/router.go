package apiserver

import "github.com/gorilla/handlers"

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	// /userproduct
	userProductRoute := s.router.PathPrefix("/userproduct").Subrouter()
	// userProductRoute.Use(s.authenticateUser)
	userProductRoute.HandleFunc("/create", s.handleUserProductCreate()).Methods("POST")

	// /product
	productRoute := s.router.PathPrefix("/product").Subrouter()
	productAuthRoute := productRoute.PathPrefix("/private").Subrouter()
	//TODO : make private and admin only
	// productAuthRoute.Use(s.authenticateUser)
	// После регистарции
	productAuthRoute.HandleFunc("/create", s.handleProductCreate()).Methods("POST")
	productAuthRoute.HandleFunc("/search", s.handleProductSearch()).Methods("POST")

	// /user
	userRoute := s.router.PathPrefix("/user").Subrouter()
	userRoute.HandleFunc("/create", s.handleUsersCreate()).Methods("POST")
	userRoute.HandleFunc("/auth", s.handleSessionsCreate()).Methods("POST")
	// После регистарции
	userAuthRoute := userRoute.PathPrefix("/private").Subrouter()
	// userAuthRoute.Use(s.authenticateUser)
	userAuthRoute.HandleFunc("/get", s.handleGetUserNow()).Methods("GET")
	userAuthRoute.HandleFunc("/delete/{id}", s.handleDeleteUser()).Methods("DELETE")
	userAuthRoute.HandleFunc("/edit/{id}", s.handleEditUser()).Methods("PUT")

	// /admin
	adminRoute := s.router.PathPrefix("/admin").Subrouter()
	adminRoute.HandleFunc("/loadproducts", s.handleLoadData()).Methods("GET")
}
