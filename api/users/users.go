package users

import (
	"database/sql"

	"github.com/kataras/iris"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token"`
}

func Register(app *iris.Framework, db *sql.DB) {
	app.Handle("POST", "/users/login", &login{db})
	app.Handle("POST", "/users/register", &register{db})
}
