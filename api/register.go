package api

import (
	"database/sql"

	"github.com/kataras/iris"
)

// 'Register' readies the middleware and mounts all package's api on an iris instance.
func Register(app *iris.Framework, db *sql.DB) {
	// Apply middleware first!
	applyMiddleware(app)

	app.Get("/", indexHandler)
	app.Handle("GET", "/todos", &todoGet{db})
}

func indexHandler(c *iris.Context) {
	c.Write("This is an example server")
}
