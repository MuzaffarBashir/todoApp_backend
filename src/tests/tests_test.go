package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
func TestApiCreateToDoSuccess(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	expectedRows := []string{"1"}
	lastUserId, _ := strconv.Atoi(expectedRows[0])
	expectedRS := sqlmock.NewRows(expectedRows).AddRow(lastUserId)
	mock.ExpectQuery(`INSERT INTO "todolist" ("description") VALUES ($1) RETURNING ID`).
		WithArgs("New Todo").
		WillReturnRows(expectedRS)

	bodydata := map[string]interface{}{
		"Description": "New Todo",
	}
	body, _ := json.Marshal(bodydata)
	response, _ := http.Post("http://localhost:8090/handlerequest", "application/json", bytes.NewReader(body))
	respBytes, _ := ioutil.ReadAll(response.Body)

	var newtodo todoapi.ApiToDo
	err := json.Unmarshal(respBytes, &newtodo)

	assert.Nil(t, err)
	assert.NotNil(t, newtodo)
	assert.EqualValues(t, "New Todo", newtodo.Description)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}
