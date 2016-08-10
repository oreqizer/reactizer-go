package server

import (
	"log"
	"net/http"
)

type Server struct {
	Mux *http.ServeMux
}

func (s *Server) MountFunc(path string, fn http.HandlerFunc) {
	s.Mux.HandleFunc(path, logger(fn))
}

func (s *Server) MountHandler(path string, handler http.Handler) {
	s.MountFunc(path, handler.ServeHTTP)
}

func NewServer() *Server {
	return &Server{ Mux: http.NewServeMux() }
}

func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[server] request received", r.URL.Path)
		fn(w, r)
		log.Println("[server] response ended")
	}
}