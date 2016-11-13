package api

import (
	"github.com/golang/glog"
	"github.com/kataras/iris"
)

func applyMiddleware(api *iris.Framework) {
	api.UseFunc(logger)
}

func logger(c *iris.Context) {
	glog.Error("Request recieved: ", c.PathString())
	c.Next()
}
