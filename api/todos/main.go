package todos

import (
	"database/sql"

	"github.com/kataras/iris"
)

type todo struct {
	Id int			`json:"id"`
	UserId int	`json:"-"`
	Text string	`json:"text"`
	Done bool		`json:"done"`
}

type todos []*todo

func Register(app *iris.Framework, db *sql.DB) {
	app.Handle("GET", "/todos", &get{db})
}
