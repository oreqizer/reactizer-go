package api

import (
	"database/sql"

	"github.com/kataras/iris"

	"reactizer-go/api/todos"
	"reactizer-go/api/users"
)

// 'Register' readies the middleware and mounts all package's api on an iris instance.
func Register(app *iris.Application, db *sql.DB) {
	// apply middleware first!
	applyMiddleware(app)

	app.Get("/", indexHandler)
	todos.Register(app.Party("/todos"), db)
	users.Register(app.Party("/users"), db)
}

func indexHandler(c iris.Context) {
	c.WriteString("This is an example server\n")
}
