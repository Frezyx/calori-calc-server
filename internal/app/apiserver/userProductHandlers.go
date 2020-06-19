package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/gorilla/mux"
)

func (s *server) handleUserProductCreate() http.HandlerFunc {
	type request struct {
		ProductID  int     `json:"productid"`
		Name       string  `json:"name"`
		Category   string  `json:"category"`
		Calory     float64 `json:"calory"`
		Squi       float64 `json:"squi"`
		Fat        float64 `json:"fat"`
		Carboh     float64 `json:"carboh"`
		Grams      float64 `json:"grams"`
		DateCreate int     `json:"date_created"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		uP := &model.UserProduct{
			ProductID:  req.ProductID,
			Name:       req.Name,
			Category:   req.Category,
			Calory:     req.Calory,
			Squi:       req.Squi,
			Carboh:     req.Carboh,
			Fat:        req.Fat,
			Grams:      req.Grams,
			DateCreate: req.DateCreate,
		}

		if err := s.store.UserProduct().Create(uP); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, uP)
	}
}

func (s *server) handleUserProductGet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		stringID := vars["id"]
		id, err := strconv.Atoi(stringID)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errNotFoundUserProduct)
			return
		}

		uP, err := s.store.UserProduct().Get(id)
		if err != nil || uP == nil {
			s.error(w, r, http.StatusNotFound, errNotFoundUserProduct)
			return
		}

		s.respond(w, r, http.StatusOK, uP)
	}
}
