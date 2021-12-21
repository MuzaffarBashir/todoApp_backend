package todoserver

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoApp/todoapi"
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

//Success case for gettig todoslist
func TestGetAllToDoSuccess(t *testing.T) {

	db := tododb.GetConnection()
	todoserver := &ToDoServer{
		Connection: db,
	}

	w := httptest.NewRecorder()
	todoserver.GetAllToDo(w, todoserver.Connection)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	var todoslist []todoapi.ApiToDo
	err := json.Unmarshal(data, &todoslist)

	assert.Nil(t, err)
	assert.NotNil(t, todoslist)
	assert.EqualValues(t, "new todo", todoslist[0].Description)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
}

//Fail case for gettig todoslist
func TestGetAllToDoFailure(t *testing.T) {

	db, _ := NewMock()
	tododb.GetConnection()
	todoserver := &ToDoServer{
		Connection: db,
	}
	w := httptest.NewRecorder()
	todoserver.GetAllToDo(w, todoserver.Connection)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	var todoslist []todoapi.ApiToDo
	_ = json.Unmarshal(data, &todoslist)
	assert.Nil(t, todoslist)
	assert.NotEqualValues(t, http.StatusOK, res.StatusCode)
}

//Success case for Creating todo
func TestCreateToDo(t *testing.T) {

	db := tododb.GetConnection()
	todoserver := &ToDoServer{
		Connection: db,
	}
	values := map[string]string{"Description": "new todo"}
	json_data, _ := json.Marshal(values)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8090/handlerequest", bytes.NewBuffer(json_data))
	response := httptest.NewRecorder()
	//calling method for creating todo
	todoserver.CreateToDo(request, response, todoserver.Connection)

	res := response.Result()

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	var todo todoapi.ApiToDo

	_ = json.Unmarshal(data, &todo)

	assert.Nil(t, err)
	assert.NotNil(t, todo)
	assert.EqualValues(t, "new todo", todo.Description)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
}

//Success case for fail Creating todo
func TestCreateToDoFail(t *testing.T) {

	db, _ := NewMock()
	todoserver := &ToDoServer{
		Connection: db,
	}
	values := map[string]string{"Description": "new todo"}
	json_data, _ := json.Marshal(values)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8090/handlerequest", bytes.NewBuffer(json_data))
	response := httptest.NewRecorder()
	//calling method for creating todo
	todoserver.CreateToDo(request, response, todoserver.Connection)

	res := response.Result()

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	var todo todoapi.ApiToDo

	_ = json.Unmarshal(data, &todo)

	assert.Nil(t, err)
	assert.NotEqualValues(t, "new todo", todo.Description)
	assert.NotEqualValues(t, http.StatusOK, res.StatusCode)
}
