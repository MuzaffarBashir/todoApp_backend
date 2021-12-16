package todoapi

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"todoApp/tododb"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	tododb.SetConnection(db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

// unit case of getting all todos
func TestGetAllToDo(t *testing.T) {

	db := tododb.GetConnection()
	todo := &ApiToDo{

		Connection: db,
	}
	todosList, err := todo.GetAllToDo(todo.Connection)
	assert.Nil(t, err)
	assert.NotNil(t, todosList)
}

// Failure unit case of getting all todos using mock db connection
func TestGetAllToDoFail(t *testing.T) {

	db, _ := NewMock()
	todo := &ApiToDo{

		Connection: db,
	}
	todosList, err := todo.GetAllToDo(db)
	assert.Nil(t, todosList)
	assert.NotNil(t, err)

}

// success case of APIcall
func TestCreateToDoApiSuccess(t *testing.T) {

	db := tododb.GetConnection()
	todo := &ApiToDo{
		Connection: db,
	}
	defer db.Close()
	bodydata := map[string]interface{}{
		"Description": "New Todo",
	}
	body, _ := json.Marshal(bodydata)
	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8090/handlerequest",
		bytes.NewReader(body))
	request.Header.Set("content-type", "text/plain")

	todo, err := todo.CreateToDoApi(request, todo.Connection)
	assert.Nil(t, err)
	assert.NotNil(t, todo)
	assert.NotNil(t, todo.ID)
	assert.EqualValues(t, "New Todo", todo.Description)
}

//success case for validation method
func TestValidateSuccess(t *testing.T) {
	description := "New Todo"
	todo := &ApiToDo{}
	newtodo, err := todo.Validate(description)
	assert.Nil(t, err)
	assert.NotNil(t, newtodo)
	assert.EqualValues(t, description, newtodo.Description)
}
