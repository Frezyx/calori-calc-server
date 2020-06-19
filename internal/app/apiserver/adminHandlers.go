package apiserver

import (
	"net/http"
)

func (s *server) handleLoadData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := "http://localhost/calory-calc-server/local/calory-calc-parser/loader.php"
		_, err := http.Get(url)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		s.respond(w, r, http.StatusCreated, "all loaded")
	}
}
