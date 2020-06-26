package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
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
