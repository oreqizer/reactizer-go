package todos

import (
	"database/sql"
	"encoding/json"

	"github.com/kataras/golog"
	"github.com/kataras/iris"

	"reactizer-go/api/utils"
)

type list struct {
	db *sql.DB
}

type create struct {
	db *sql.DB
}

type edit struct {
	db *sql.DB
}

type remove struct {
	db *sql.DB
}

// You can also do something like that if you have just one dependency:
// func ListHandler(db *sql.DB) iris.Handler{
// 	return func(c iris.Context){
//		// [...]
// 	}
// }
func (r *list) Serve(c iris.Context) {
	uid, err := utils.Authorize(c)
	if err != nil {
		utils.Error(c, err.Error(), 401)
		return
	}

	rows, err := r.db.Query("SELECT * FROM todos WHERE user_id=$1", uid)
	if err != nil {
		golog.Error(err)
		return
	}
	defer rows.Close()

	todos, err := scanTodos(rows)
	if err != nil {
		golog.Error(err)
		return
	}

	data, err := json.Marshal(todos)
	if err != nil {
		golog.Error(err)
		return
	}

	c.Write(data)
}

func (r *create) Serve(c iris.Context) {
	uid, err := utils.Authorize(c)
	if err != nil {
		utils.Error(c, err.Error(), 401)
		return
	}

	todo := &Todo{UserId: uid}
	err = c.ReadJSON(todo)
	if err != nil {
		golog.Error(err)
		return
	}
	err = r.db.QueryRow(`
		INSERT INTO todos (text, user_id, done)
		VALUES ($1, $2, $3) RETURNING id
	`, todo.Text, todo.UserId, false).Scan(&todo.Id)
	if err != nil {
		golog.Error(err)
		return
	}

	c.JSON(todo)
}

func (r *edit) Serve(c iris.Context) {
	uid, err := utils.Authorize(c)
	if err != nil {
		utils.Error(c, err.Error(), 401)
		return
	}

	// uint64 and GetUint64 is better for IDs*
	id, err := c.Params().GetInt("id")
	if err != nil {
		golog.Error(err)
		return
	}
	todo := &Todo{Id: id, UserId: uid}
	err = c.ReadJSON(todo)
	if err != nil {
		golog.Error(err)
		return
	}
	res, err := r.db.Exec(`
		UPDATE todos SET text=$1, done=$2 WHERE id=$3 AND user_id=$4
	`, todo.Text, todo.Done, todo.Id, todo.UserId)
	if err != nil {
		golog.Error(err)
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		golog.Error(err)
		return
	}
	if count == 0 {
		utils.Error(c, notFound, 404)
		return
	}

	c.JSON(todo)
}

func (r *remove) Serve(c iris.Context) {
	uid, err := utils.Authorize(c)
	if err != nil {
		utils.Error(c, err.Error(), 401)
		return
	}

	id, err := c.Params().GetInt("id")
	if err != nil {
		golog.Error(err)
		return
	}
	res, err := r.db.Exec("DELETE FROM todos WHERE id=$1 AND user_id=$2", id, uid)
	if err != nil {
		golog.Error(err)
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		golog.Error(err)
		return
	}
	if count == 0 {
		utils.Error(c, notFound, 404)
		return
	}

	c.Text("ok")
}
