package todoapi

import (
	"database/sql"
	"todoApp/tododb"
)

func (apiTodo *ApiToDo) GetAllToDo(conn *sql.DB) ([]ApiToDo, error) {

	var todosList []ApiToDo
	tododb := tododb.NewTodo()
	data, err := tododb.GetToDO(conn)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		newtodo := NewApiService()
		if err := data.Scan(&newtodo.ID, &newtodo.Description); err != nil {
			panic(err)
		}
		todosList = append(todosList, *newtodo)
	}
	return todosList, nil
}
