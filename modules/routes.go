package modules

import (
	"net/http"
)

type Mountable interface {
	MountFunc(path string, fn http.HandlerFunc)
	MountHandler(path string, fn http.Handler)
}

func MountRoutes(mux Mountable) {
	mux.MountFunc("/", indexHandler)
	mux.MountFunc("/todos", todoHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}