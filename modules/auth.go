package modules

import (
	"database/sql"
	"net/http"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

// 'authorize' checks the 'X-Authorization' header if it contains the JWT token required by some
// queries. If the token is there, it is decoded into a user id and returned.
//
// In case of an error, translation id AuthError is returned.
func authorize(r *http.Request, db *sql.DB) (string, error) {
	// TODO make generic interface/struct for T
	token := r.Header["X-Authorization"]
	if len(token) != 1 {
		return "", AuthError("auth.no_auth_header")
	}
	return "", nil
}
