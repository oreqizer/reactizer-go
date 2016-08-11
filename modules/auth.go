package modules

import (
	"database/sql"
	"net/http"
	"github.com/nicksnyder/go-i18n/i18n"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

func authorize(r *http.Request, db *sql.DB, T i18n.TranslateFunc) (int, error) {
	// TODO make generic interface/struct for T
	token := r.Header["X-Authorization"]
	if len(token) != 1 {
		return 0, AuthError(T("auth.no_auth_header"))
	}
	return 0, nil
}
