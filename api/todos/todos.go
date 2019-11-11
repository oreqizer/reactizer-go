package todos

import (
	"database/sql"

	"github.com/kataras/iris/v12"
)

type Todo struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
}

// Check
// https://github.com/kataras/iris/tree/master/_examples/README.md#mvc
// to learn how you can convert all these to controllers, I let that for you, as an exercise.
func Register(r iris.Party, db *sql.DB) {
	var (
		listCtrl   = &list{db}
		createCtrl = &create{db}
		editCtrl   = &edit{db}
		removeCtrl = &remove{db}
	)

	r.Get("/", listCtrl.Serve)
	r.Post("/", createCtrl.Serve)
	r.Put("/{id:int}", editCtrl.Serve) // {id:uint64} will be better*
	r.Delete("/{id:int}", removeCtrl.Serve)
}
