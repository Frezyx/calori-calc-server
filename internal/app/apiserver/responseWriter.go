package apiserver

import "net/http"

type responseCustomWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseCustomWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
