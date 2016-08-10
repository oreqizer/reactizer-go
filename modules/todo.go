package modules

import (
	"net/http"
)

type Todo struct {
	Id int
	UserId int
	Text string
	Done bool
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is a todo.\n"))
}
