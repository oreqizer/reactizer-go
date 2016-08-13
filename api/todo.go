package api

import (
	"database/sql"
	"log"
	"encoding/json"

	"github.com/kataras/iris"
)

type Todo struct {
	Id int			`json:"id"`
	UserId int	`json:"-"`
	Text string	`json:"text"`
	Done bool		`json:"done"`
}

type Todos []*Todo

type todoGet struct {
	db *sql.DB
}

func (t *todoGet) Serve(c *iris.Context) {
	T := getT(c)
	_, err := authorize(c, t.db)
	if err != nil {
		c.SetStatusCode(401)
		c.Write(T(err.Error()))
		return
	}

	rows, err := t.db.Query("SELECT * FROM todo")
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()

	todos, err := scanTodos(rows)
	if err != nil {
		log.Print(err)
		return
	}

	data, err := json.Marshal(todos)
	if err != nil {
		log.Print(err)
		return
	}

	c.Write(string(data))
}

func scanTodos(rows *sql.Rows) (Todos, error) {
	todos := Todos{}
	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Text, &todo.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}