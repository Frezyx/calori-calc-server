package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
)

func (s *server) handleProductCreate() http.HandlerFunc {
	type request struct {
		Name     string  `json:"name"`
		Category string  `json:"category"`
		Calory   float64 `json:"calory"`
		Squi     float64 `json:"squi"`
		Fat      float64 `json:"fat"`
		Carboh   float64 `json:"carboh"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		p := &model.Product{
			Name:     req.Name,
			Category: req.Category,
			Calory:   req.Calory,
			Squi:     req.Squi,
			Carboh:   req.Carboh,
			Fat:      req.Fat,
		}

		if err := s.store.Product().Create(p); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, p)
	}
}
