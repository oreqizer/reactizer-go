package server

import (
	"log"
	"net/http"
)

type ServeBox struct {
	Mux *http.ServeMux
}

func (s *ServeBox) HandleFunc(path string, fn http.HandlerFunc) {
	s.Mux.HandleFunc(path, logger(fn))
}

func (s *ServeBox) Handle(path string, handler http.Handler) {
	s.HandleFunc(path, handler.ServeHTTP)
}

func NewServeBox() *ServeBox {
	return &ServeBox{ Mux: http.NewServeMux() }
}

func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[server] request received", r.URL.Path)
		fn(w, r)
	}
}