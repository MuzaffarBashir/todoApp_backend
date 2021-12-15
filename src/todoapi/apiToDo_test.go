package todoapi

import (
	"database/sql"
	"log"
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

	db, _ := NewMock()
	todo := &ApiToDo{

		Connection: db,
	}
	todosList, err := todo.GetAllToDo(todo.Connection)
	assert.Nil(t, err)
	assert.NotNil(t, todosList)

}
