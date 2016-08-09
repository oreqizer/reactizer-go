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

func init() {
	http.HandleFunc("/api/todos", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {

}