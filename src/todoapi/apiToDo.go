package todoapi

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
func (apitodo *ApiToDo) CreateToDoApi(request *http.Request, conn *sql.DB) (*ApiToDo, error) {

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&apitodo)
	if err != nil {
		err = fmt.Errorf("bad Request")
		return nil, err
	}
	// validating todo description for db model
	tododb, err := apitodo.Validate(apitodo.Description)
	if err != nil {
		err = errors.New("description missing Bad Request")
		return nil, err
	} else {

		ID, err := tododb.CreateToDo(conn)
		if err != nil {
			err = errors.New("database internal error")
			return nil, err
		}
		idString := strconv.Itoa(int(ID))
		apitodo.ID = idString
	}
	return apitodo, nil
}
func (apitodo *ApiToDo) Validate(description string) (*tododb.Todo, error) {

	tododb := tododb.NewTodo() // cretaing db model for todo insertion
	if len(strings.TrimSpace(description)) > 0 {
		tododb.Description = description
		return tododb, nil
	}
	return nil, errors.New("please enter description")
}
