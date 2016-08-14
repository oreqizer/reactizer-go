package api

import (
	"database/sql"

	"github.com/kataras/iris"

	"reactizer-go/api/todos"
	"reactizer-go/api/users"
)

// 'Register' readies the middleware and mounts all package's api on an iris instance.
func Register(app *iris.Framework, db *sql.DB) {
	// apply middleware first!
	applyMiddleware(app)

	app.Get("/", indexHandler)
	todos.Register(app, db)
	users.Register(app, db)
}

func indexHandler(c *iris.Context) {
	c.Write("This is an example server")
}
