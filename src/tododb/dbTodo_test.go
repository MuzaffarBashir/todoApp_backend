package tododb

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	SetConnection(db)
	if err != nil {
		log.Fatalf("wasn't thinking of err '%s' while geting connection", err)
	}
	return db, mock
}

//success case for get todos from db
func TestGetAllToDO(t *testing.T) {
	db, mock := NewMock()
	todo := &Todo{
		Connection: db,
	}
	query := `SELECT id,description FROM "todolist"`
	expectedRows := []string{"id", "description"}
	expectedRS := sqlmock.NewRows(expectedRows).FromCSVString("1 new todo")
	mock.ExpectQuery(query).WillReturnRows(expectedRS)
	actualTodo, err := todo.GetAllToDO(todo.Connection)
	totalActualRows := 0
	for actualTodo.Next() {
		totalActualRows++
	}
	assert.Nil(t, err)
	assert.NotNil(t, actualTodo)
	assert.EqualValues(t, 1, totalActualRows)
}

//Failure case for geting list of  todo
func TestGetToDO_EmptyTodoList(t *testing.T) {
	db, mock := NewMock()
	todo := &Todo{
		Connection: db,
	}
	defer db.Close()
	query := `SELECT id,description FROM "todolist"`
	expectedRows := []string{}
	expectedRS := sqlmock.NewRows(expectedRows)
	mock.ExpectQuery(query).WillReturnRows(expectedRS)
	actualTodo, actualErr := todo.GetAllToDO(todo.Connection)
	totalActualRows := 0
	for actualTodo.Next() {
		totalActualRows++
	}
	assert.Nil(t, actualErr)
	assert.NotNil(t, actualTodo)
	assert.Equal(t, len(expectedRows), totalActualRows)
}
