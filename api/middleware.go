package api

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

func applyMiddleware(api *iris.Application) {
	api.Use(logger)
}

func logger(c iris.Context) {
	golog.Infof("Request recieved: %s", c.Path())
	c.Next()
}
