package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/gorilla/mux"
)

func (s *server) handleDateCreate() http.HandlerFunc {
	type request struct {
		IDs  string `json:"products_ids"`
		Date int    `json:"date_created"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		d := &model.Date{
			IDs:  req.IDs,
			Date: req.Date,
		}

		if err := s.store.Dates().Create(d); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, d)
	}
}

func (s *server) handleDateIsSet() http.HandlerFunc {

	type request struct {
		Date int `json:"date_created"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		dateID, err := s.store.Dates().GetIfSet(req.Date)
		if err != nil {
			s.respond(w, r, http.StatusNotFound, dateID)
		}
		s.respond(w, r, http.StatusOK, dateID)

	}
}

func (s *server) handleGetIDsByDate() http.HandlerFunc {

	type request struct {
		Date int `json:"date_created"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		dateIDs, err := s.store.Dates().GetIDsByDate(req.Date)
		if err != nil {
			s.respond(w, r, http.StatusNotFound, dateIDs)
		}
		s.respond(w, r, http.StatusOK, dateIDs)

	}
}

func (s *server) updateDietHandler() http.HandlerFunc {

	type request struct {
		Date int    `json:"date_created"`
		IDs  string `json:"products_ids"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		vars := mux.Vars(r)
		stringID := vars["id"]
		idMux, err := strconv.Atoi(stringID)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errNotFoundDate)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		d := &model.Date{
			ID:   idMux,
			IDs:  req.IDs,
			Date: req.Date,
		}

		if err := s.store.Dates().UpdateDate(d); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, msgChangesSave)

	}
}
