package todos

import (
	"database/sql"

	"github.com/kataras/iris"
)

type Todo struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
}

func Register(app *iris.Framework, db *sql.DB) {
	app.Handle("GET", "/todos", &list{db})
	app.Handle("POST", "/todos", &create{db})
	app.Handle("PUT", "/todos/:id", &edit{db})
	app.Handle("DELETE", "/todos/:id", &remove{db})
}
