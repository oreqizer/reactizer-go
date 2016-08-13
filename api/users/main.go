package users

import (
	"database/sql"

	"github.com/kataras/iris"
)

type user struct {
	Id int						`json:"id"`
	Username string		`json:"username"`
	Email string			`json:"email"`
	Password string		`json:"password"`
}

func Register(app *iris.Framework, db *sql.DB) {

}