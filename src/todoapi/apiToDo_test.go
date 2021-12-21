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

//success case for validation method
func TestValidateSuccess(t *testing.T) {
	description := "New Todo"
	todo := &ApiToDo{}
	newtodo, err := todo.Validate(description)
	assert.Nil(t, err)
	assert.NotNil(t, newtodo)
	assert.EqualValues(t, description, newtodo.Description)
}

//Failure case for validate method
func TestValidateFailure(t *testing.T) {

	description := ""
	todo := &ApiToDo{}
	newtodo, err := todo.Validate(description)
	assert.Nil(t, newtodo)
	assert.NotNil(t, err)
}
