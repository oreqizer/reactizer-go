package api

import (
	"github.com/kataras/iris"
	"github.com/golang/glog"
)

func applyMiddleware(api *iris.Framework) {
	api.UseFunc(logger)
}

func logger(c *iris.Context) {
	glog.Error("[server] request recieved: ", c.PathString())
	c.Next()
}
