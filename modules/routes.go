package modules

import (
	"net/http"
	"database/sql"

	"github.com/nicksnyder/go-i18n/i18n"
	"reactizer-go/config"
)

type Mountable interface {
	MountFunc(path string, fn http.HandlerFunc)
	MountHandler(path string, fn http.Handler)
}

func Register(mux Mountable, db *sql.DB) {
	todos := &todoHandler{db: db}

	mux.MountFunc("/", indexHandler)
	mux.MountHandler("/todos", todos)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func getT(r *http.Request) (i18n.TranslateFunc, error) {
	acceptLang := r.Header.Get("Accept-Language")
	T, err := i18n.Tfunc(acceptLang, config.DefaultLanguage)
	if err != nil {
		return nil, err
	}
	return T, nil
}
