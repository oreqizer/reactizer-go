package todos

import "database/sql"

// 'scanTodos' scans all the rows of the query and returns them as a slice of Todos.
func scanTodos(rows *sql.Rows) ([]todo, error) {
	todos := []todo{}
	for rows.Next() {
		todo := &todo{}
		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Text, &todo.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
