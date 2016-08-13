package api

import (
	"database/sql"

	"github.com/kataras/iris"

	"reactizer-go/api/todos"
)

// 'Register' readies the middleware and mounts all package's api on an iris instance.
func Register(app *iris.Framework, db *sql.DB) {
	// Apply middleware first!
	applyMiddleware(app)

	app.Get("/", indexHandler)
	todos.Register(app, db)
}

func indexHandler(c *iris.Context) {
	c.Write("This is an example server")
}
