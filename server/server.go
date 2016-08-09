package server

import (
	"log"
	"net/http"
)

var mux *http.ServeMux = http.NewServeMux()

func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[server] request received", r.URL.Path)
		fn(w, r)
		log.Println("[server] response ended")
	}
}

func Mount(path string, fn http.HandlerFunc) {
	mux.HandleFunc(path, logger(fn))
}

func MountHandler(path string, handler http.Handler) {
	Mount(path, handler.ServeHTTP)
}

func Start(addr string) {
	log.Fatal(http.ListenAndServe(addr, mux))
}
