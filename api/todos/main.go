package todos

import (
	"database/sql"

	"github.com/kataras/iris"
)

type Todo struct {
	Id int			`json:"id"`
	UserId int	`json:"-"`
	Text string	`json:"text"`
	Done bool		`json:"done"`
}

func Register(app *iris.Framework, db *sql.DB) {
	app.Handle("GET", "/todos", &get{db})
}
