package modules

import (
	"net/http"
	"database/sql"
)

type Handlable interface {
	HandleFunc(path string, fn http.HandlerFunc)
	Handle(path string, fn http.Handler)
}

// 'Register' readies the handlers and mounts all package's routes on a given multiplexor.
func Register(mux Handlable, db *sql.DB) {
	todos := &todoHandler{db: db}

	mux.HandleFunc("/", indexHandler)
	mux.Handle("/todos", todos)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

