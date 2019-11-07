package users

import (
	"database/sql"

	"github.com/kataras/iris/v12"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token"`
}

func Register(r iris.Party, db *sql.DB) {
	var (
		loginCtrl    = &login{db}
		registerCtrl = &register{db}
	)

	r.Post("/login", loginCtrl.Serve)
	r.Post("/register", registerCtrl.Serve)
}
