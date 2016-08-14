package todos

import (
	"database/sql"
	"encoding/json"

	"github.com/kataras/iris"
	"github.com/golang/glog"

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

func (r *list) Serve(c *iris.Context) {
	T := utils.GetT(c)
	uid, err := utils.Authorize(c)
	if err != nil {
		c.Error(T(err.Error()), 401)
		return
	}

	rows, err := r.db.Query("SELECT * FROM todos WHERE user_id=$1", uid)
	if err != nil {
		glog.Error(err)
		return
	}
	defer rows.Close()

	todos, err := scanTodos(rows)
	if err != nil {
		glog.Error(err)
		return
	}

	data, err := json.Marshal(todos)
	if err != nil {
		glog.Error(err)
		return
	}

	c.Write(string(data))
}

func (r *create) Serve(c *iris.Context) {
	T := utils.GetT(c)
	uid, err := utils.Authorize(c)
	if err != nil {
		c.Error(T(err.Error()), 401)
		return
	}

	todo := &Todo{UserId: uid}
	err = c.ReadJSON(todo)
	if err != nil {
		glog.Error(err)
		return
	}
	err = r.db.QueryRow(`
		INSERT INTO todos (text, user_id, done)
		VALUES ($1, $2, $3) RETURNING id
	`, todo.Text, todo.UserId, false).Scan(&todo.Id)
	if err != nil {
		glog.Error(err)
		return
	}

	c.JSON(200, todo)
}

func (r *edit) Serve(c *iris.Context) {
	T := utils.GetT(c)
	uid, err := utils.Authorize(c)
	if err != nil {
		c.Error(T(err.Error()), 401)
		return
	}

	id, err := c.ParamInt("id")
	if err != nil {
		glog.Error(err)
		return
	}
	todo := &Todo{Id: id, UserId: uid}
	err = c.ReadJSON(todo)
	if err != nil {
		glog.Error(err)
		return
	}
	res, err := r.db.Exec(`
		UPDATE todos SET text=$1, done=$2 WHERE id=$3 AND user_id=$4
	`, todo.Text, todo.Done, todo.Id, todo.UserId)
	if err != nil {
		glog.Error(err)
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		glog.Error(err)
		return
	}
	if count == 0 {
		c.Error(T(notFound), 404)
		return
	}

	c.JSON(200, todo)
}

func (r *remove) Serve(c *iris.Context) {
	T := utils.GetT(c)
	uid, err := utils.Authorize(c)
	if err != nil {
		c.Error(T(err.Error()), 401)
		return
	}

	id, err := c.ParamInt("id")
	if err != nil {
		glog.Error(err)
		return
	}
	res, err := r.db.Exec("DELETE FROM todos WHERE id=$1 AND user_id=$2", id, uid)
	if err != nil {
		glog.Error(err)
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		glog.Error(err)
		return
	}
	if count == 0 {
		c.Error(T(notFound), 404)
		return
	}

	c.Text(200, "ok")
}