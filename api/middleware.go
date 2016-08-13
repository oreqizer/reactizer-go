package api

import (
	"log"

	"github.com/kataras/iris"
)

func applyMiddleware(api *iris.Framework) {
	api.UseFunc(logger)
}

func logger(c *iris.Context) {
	log.Print("[server] request recieved: ", c.PathString())
	c.Next()
}
