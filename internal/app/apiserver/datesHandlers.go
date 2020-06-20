package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
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
