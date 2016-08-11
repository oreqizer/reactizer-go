package modules

import (
	"database/sql"
	"net/http"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

func authorize(r *http.Request, db *sql.DB) (int, error) {
	token := r.Header["X-Authorization"]
	if len(token) != 1 {
		return 0, AuthError("no auth header")
	}
	return 0, nil
}
