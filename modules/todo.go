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

type todoHandler struct {
	db *sql.DB
}

func (t *todoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rows, err := t.db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	todos := []*Todo{}
	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Text, &todo.Done)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	json, err := json.Marshal(todos)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
