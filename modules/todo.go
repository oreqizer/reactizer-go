package modules

import (
	"net/http"
	"database/sql"
	"log"
	"encoding/json"
)

type Todo struct {
	Id int			`json:"id"`
	UserId int	`json:"-"`
	Text string	`json:"text"`
	Done bool		`json:"done"`
}

type Todos []*Todo

type todoHandler struct {
	db *sql.DB
}

func (t *todoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	T := getT(r)
	_, err := authorize(r, t.db)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(T(err.Error())))
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

	json, err := json.Marshal(todos)
	if err != nil {
		log.Print(err)
		return
	}

	w.Write(json)
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