package todo

import (
	"net/http"
	"reactizer-go/server"
)

type Todo struct {
	Id int
	UserId int
	Text string
	Done bool
}

func Mount() {
	server.Mount("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}
