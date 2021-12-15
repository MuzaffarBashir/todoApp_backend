package todoserver

import (
	"database/sql"
	"log"
	"net/http"
	"testing"
	"todoApp/todoapi"
	"todoApp/tododb"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// Success case of getting all todos
func TestGetTODO(t *testing.T) {

	db, _ := NewMock()
	todo := &todoapi.ApiToDo{

		Connection: db,
	}
	var response http.ResponseWriter
	response, err := todo.GetTodo(todo.Connection)
	assert.Nil(t, err)
	assert.NotNil(t, response)

}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	tododb.SetConnection(db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
