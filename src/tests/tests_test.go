package tests

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

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

	query := `SELECT id,description FROM "todolist"`
	expectedRows := []string{"id", "description"}
	expectedRS := sqlmock.NewRows(expectedRows).FromCSVString("1 New TODO")
	mock.ExpectQuery(query).WillReturnRows(expectedRS)
	response, _ := http.Get("http://localhost:8090/gettodo")
	bytes, _ := ioutil.ReadAll(response.Body)
	var todoslist = make([]todoapi.ApiToDo, 0)
	err := json.Unmarshal(bytes, &todoslist)
	assert.Nil(t, err)
	assert.NotNil(t, todoslist)
	assert.EqualValues(t, "new todo", todoslist[0].Description)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}
