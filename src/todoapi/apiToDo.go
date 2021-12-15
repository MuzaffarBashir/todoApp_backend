package todoapi

import (
	"database/sql"
	"fmt"
	"todoApp/tododb"
)

func (apiTodo *ApiToDo) GetAllToDo(conn *sql.DB) ([]ApiToDo, error) {

	var todosList []ApiToDo
	tododb := tododb.NewTodo()
	data, err := tododb.GetAllToDO(conn)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		newtodo := NewApiService()
		if err := data.Scan(&newtodo.ID, &newtodo.Description); err != nil {
			fmt.Print("Error in apitodo")
			panic(err)
		}
		todosList = append(todosList, *newtodo)
	}
	return todosList, nil
}
