package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/gorilla/mux"
)

func (s *server) handleDietCreate() http.HandlerFunc {
	type request struct {
		Name          string `json:"name"`
		UserID        int    `json:"user_id"`
		IsAutoCreated bool   `json:"is_auto_created"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			ID: req.UserID,
		}

		if err := s.store.Diets().Create(u, req.Name, req.IsAutoCreated); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, msgDietCreate)
	}
}

func (s *server) handleDietGet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		stringID := vars["id"]
		id, err := strconv.Atoi(stringID)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		d, err := s.store.Diets().GetByID(id)
		if err != nil || d == nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, d)
	}
}

func (s *server) handleAllDietGetByUserID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		stringID := vars["userID"]
		userID, err := strconv.Atoi(stringID)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		dList, err := s.store.Diets().GetAllByUserID(userID)
		if err != nil || dList == nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, dList)
	}
}
