package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
func TestApiGetAllToDos(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	expectedRS := sqlmock.NewRows([]string{"id", "description"}).FromCSVString("1, new todo")
	mock.ExpectQuery("SELECT (.*) FROM todolist").WillReturnRows(expectedRS)

	rs, err := http.Get("http://localhost:8090/gettodo")
	bytes, _ := ioutil.ReadAll(rs.Body)
	var todoslist = make([]todoapi.ApiToDo, 0)
	err = json.Unmarshal(bytes, &todoslist)
	assert.Nil(t, err)
	assert.NotNil(t, todoslist)
	assert.EqualValues(t, "new todo", todoslist[0].Description)
	assert.EqualValues(t, http.StatusOK, rs.StatusCode)
}
func TestApiCreateToDoSuccess(t *testing.T) {

	db, mock := NewMock()
	defer db.Close()
	prep := mock.ExpectPrepare("INSERT INTO todolist (.+) VALUES (.+)")
	prep.ExpectExec().WithArgs("new todo").WillReturnResult(sqlmock.NewResult(1, 1))

	var body = []byte(`{"description": "new todo"}`)
	response, _ := http.Post("http://localhost:8090/handlerequest", "application/json", bytes.NewReader(body))
	respBytes, _ := ioutil.ReadAll(response.Body)
	var newtodo todoapi.ApiToDo
	err := json.Unmarshal(respBytes, &newtodo)

	assert.Nil(t, err)
	assert.NotNil(t, newtodo)
	assert.EqualValues(t, "new todo", newtodo.Description)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}
