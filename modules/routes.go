package modules

import (
	"net/http"
	"database/sql"
)

type Mountable interface {
	MountFunc(path string, fn http.HandlerFunc)
	MountHandler(path string, fn http.Handler)
}

// 'Register' readies the handlers and mounts all package's routes on a given multiplexor.
func Register(mux Mountable, db *sql.DB) {
	todos := &todoHandler{db: db}

	mux.MountFunc("/", indexHandler)
	mux.MountHandler("/todos", todos)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

