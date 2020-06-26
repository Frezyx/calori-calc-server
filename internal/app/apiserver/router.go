package apiserver

import "github.com/gorilla/handlers"

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	// /diet
	dietRoute := s.router.PathPrefix("/diet").Subrouter()
	dietRoute.HandleFunc("/autocreate", s.handleDietCreate()).Methods("POST")
	dietRoute.HandleFunc("/delete/{id}", s.handleDeleteDiet()).Methods("DELETE")
	dietRoute.HandleFunc("/get/{id}", s.handleDietGet()).Methods("GET")
	dietRoute.HandleFunc("/getall/{userID}", s.handleAllDietGetByUserID()).Methods("GET")
	dietRoute.HandleFunc("/update/{id}", s.handleUpdateDiet()).Methods("PUT")

	// /date
	datesRoute := s.router.PathPrefix("/date").Subrouter()
	datesRoute.HandleFunc("/create", s.handleDateCreate()).Methods("POST")
	datesRoute.HandleFunc("/{id}/update", s.updateDietHandler()).Methods("PUT")
	// /get
	datesGetRoute := datesRoute.PathPrefix("/get").Subrouter()
	datesGetRoute.HandleFunc("/bydate", s.handleGetIDsByDate()).Methods("GET")
	datesGetRoute.HandleFunc("/isset", s.handleDateIsSet()).Methods("GET")

	// /userproduct
	userProductRoute := s.router.PathPrefix("/userproduct").Subrouter()
	// userProductRoute.Use(s.authenticateUser)
	// После авторизации
	userProductRoute.HandleFunc("/create", s.handleUserProductCreate()).Methods("POST")
	userProductRoute.HandleFunc("/get/{id}", s.handleUserProductGet()).Methods("GET")
	userProductRoute.HandleFunc("/getall/{userID}", s.handleUserProductGetAll()).Methods("GET")
	userProductRoute.HandleFunc("/edit/{id}", s.handleEditUserProduct()).Methods("PUT")
	userProductRoute.HandleFunc("/delete/{id}", s.handleDeleteProduct()).Methods("DELETE")
	userProductRoute.HandleFunc("/deleteall", s.handleDeleteAllProduct()).Methods("DELETE")

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
	userAuthRoute.HandleFunc("/getbyid/{id}", s.handleUserGetByID()).Methods("GET")

	// /admin
	adminRoute := s.router.PathPrefix("/admin").Subrouter()
	adminRoute.HandleFunc("/loadproducts", s.handleLoadData()).Methods("GET")
}
