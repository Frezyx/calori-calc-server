package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/gorilla/mux"
)

func (s *server) handleGetUserNow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email:    req.Email,
			Name:     req.Name,
			Surname:  req.Surname,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		s.respond(w, r, http.StatusOK, msgAuthorized)
	}
}

func (s *server) handleDeleteUser() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		stringID := vars["id"]
		id, err := strconv.Atoi(stringID)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errNotFoundUser)
			return
		}

		cond, err := s.store.User().DeleteUser(id)
		if err != nil || !cond {
			s.error(w, r, http.StatusNotFound, errNotFoundUser)
			return
		}

		s.respond(w, r, http.StatusOK, msgUserDeleted)
	}
}

func (s *server) handleEditUser() http.HandlerFunc {

	type request struct {
		Email           string  `json:"email"`
		Name            string  `json:"name"`
		Surname         string  `json:"surname"`
		Weight          float64 `json:"weight"`
		Height          float64 `json:"height"`
		Age             int     `json:"age"`
		WorkModel       float64 `json:"workmodel"`
		WorkFutureModel float64 `json:"workfuturemodel"`
		Gender          bool    `json:"gender"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		stringID := vars["id"]
		idMux, err := strconv.Atoi(stringID)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errNotFoundUser)
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			ID:              idMux,
			Email:           req.Email,
			Name:            req.Name,
			Surname:         req.Surname,
			Weight:          req.Weight,
			Height:          req.Height,
			Age:             req.Age,
			WorkModel:       req.WorkModel,
			WorkFutureModel: req.WorkFutureModel,
			Gender:          req.Gender,
		}
		if err := s.store.User().Edit(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, msgChangesSave)
	}
}

func (s *server) handleUserGetByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		stringID := vars["id"]
		id, err := strconv.Atoi(stringID)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().GetByID(id)
		if err != nil || u == nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, u)
	}
}
